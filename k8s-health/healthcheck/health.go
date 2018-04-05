/*
 * Copyright 2018, Oath Inc.
 * Licensed under the terms of the MIT license. See LICENSE file in the project root for terms.
 */

package healthcheck

import (
    "github.com/serhii-samoilenko/pod-startup-lock/k8s-health/config"
    "github.com/serhii-samoilenko/pod-startup-lock/k8s-health/k8s"
    "github.com/serhii-samoilenko/pod-startup-lock/common/util"
    "time"
    "log"
    "k8s.io/api/extensions/v1beta1"
    "k8s.io/api/core/v1"
    "fmt"
)

type HealthChecker struct {
    k8s        *k8s.Client
    conf       config.Config
    nodeLabels map[string]string
    isHealthy  bool
}

func NewHealthChecker(appConfig config.Config, k8s *k8s.Client) *HealthChecker {
    nodeLabels := k8s.GetNodeLabels(appConfig.NodeName)
    return &HealthChecker{k8s, appConfig, nodeLabels, false}
}

func (h *HealthChecker) HealthFunction() func() bool {
    return func() bool {
        return h.isHealthy
    }
}

func (h *HealthChecker) Run() {
    for {
        if h.check() {
            log.Print("HealthCheck passed")
            h.isHealthy = true
            time.Sleep(h.conf.HealthPassTimeout)
        } else {
            log.Print("HealthCheck failed")
            h.isHealthy = false
            time.Sleep(h.conf.HealthFailTimeout)
        }
    }
}

func (h *HealthChecker) check() bool {
    log.Print("---")
    log.Print("HealthCheck:")
    daemonSets := h.k8s.GetDaemonSets(h.conf.Namespace)
    if h.checkAllDaemonSetsReady(daemonSets) {
        return true
    }
    nodePods := h.k8s.GetNodePods(h.conf.NodeName)
    return h.checkAllDaemonSetsPodsAvailableOnNode(daemonSets, nodePods)
}

func (h *HealthChecker) checkAllDaemonSetsReady(daemonSets []v1beta1.DaemonSet) bool {
    for _, ds := range daemonSets {
        if required, reason := h.checkRequired(&ds); !required {
            log.Print(reason)
            continue
        }
        status := ds.Status
        if status.DesiredNumberScheduled != status.NumberReady {
            log.Printf("'%v' daemonSet not ready: Desired: '%v', Ready: '%v'",
                ds.Name, status.DesiredNumberScheduled, status.NumberReady)
            return false
        }
        log.Printf("'%v': ok", ds.Name)
    }
    log.Print("All DaemonSets ok")
    return true
}

func (h *HealthChecker) checkAllDaemonSetsPodsAvailableOnNode(daemonSets []v1beta1.DaemonSet, pods []v1.Pod) bool {
    for _, ds := range daemonSets {
        if required, reason := h.checkRequired(&ds); !required {
            log.Print(reason)
            continue
        }
        log.Printf("'%v' daemonSet: Looking for Pods on node", ds.Name)
        pod, found := findDaemonSetPod(&ds, pods)
        if !found {
            log.Printf("'%v' daemonSet: No Pods found", ds.Name)
            return false
        }
        log.Printf("'%v' daemonSet: Found Pod: '%v'", ds.Name, pod.Name)
        if !isPodReady(pod) {
            return false
        }
    }
    log.Print("All DaemonSets Pods available on node")
    return true
}

func (h *HealthChecker) checkRequired(ds *v1beta1.DaemonSet) (bool, string) {
    reason := fmt.Sprintf("'%v' daemonSet Excluded from healthcheck: ", ds.Name)
    if len(h.conf.ExcludeDs) > 0 && util.MapContainsAnyPair(ds.Labels, h.conf.ExcludeDs) {
        return false, reason + "matches exclude labels"
    }
    if len(h.conf.IncludeDs) > 0 && !util.MapContainsAllPairs(ds.Labels, h.conf.IncludeDs) {
        return false, reason + "not matches include labels"
    }
    if h.conf.HostNetworkDs && !ds.Spec.Template.Spec.HostNetwork {
        return false, reason + "not on host network"
    }
    nodeSelector := ds.Spec.Template.Spec.NodeSelector
    if !util.MapContainsAll(h.nodeLabels, nodeSelector) {
        return false, reason + "not eligible for scheduling on node"
    }
    return true, fmt.Sprintf("'%v' daemonSet healthcheck required", ds.Name)
}

func findDaemonSetPod(ds *v1beta1.DaemonSet, pods []v1.Pod) (*v1.Pod, bool) {
    for _, pod := range pods {
        if isPodOwnedByDs(&pod, ds) {
            return &pod, true
        }
    }
    return nil, false
}

func isPodReady(pod *v1.Pod) bool {
    if pod.Status.Phase != "Running" {
        log.Printf("'%v' Pod: Not running: Phase: '%v'", pod.Name, pod.Status.Phase)
        return false
    }
    for _, cond := range pod.Status.Conditions {
        if cond.Type == "Ready" && cond.Status == "True" {
            log.Printf("'%v' Pod: Ready", pod.Name)
            return true
        }
    }
    log.Printf("'%v' Pod: Not Ready: '%v'", pod.Name, pod.Status.Conditions)
    return false
}

func isPodOwnedByDs(pod *v1.Pod, ds *v1beta1.DaemonSet) bool {
    for _, ref := range pod.ObjectMeta.OwnerReferences {
        if ds.ObjectMeta.UID == ref.UID {
            return true
        }
    }
    return false
}
