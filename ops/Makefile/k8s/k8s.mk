# Include Makefile
include $(SELF_DIR)/ops/Makefile/k8s/helm.mk
include $(SELF_DIR)/ops/Makefile/k8s/check.mk
include $(SELF_DIR)/ops/Makefile/k8s/ct.mk
include $(SELF_DIR)/ops/Makefile/k8s/k8s.velero.mk
include $(SELF_DIR)/ops/Makefile/k8s/kubeadm.mk
include $(SELF_DIR)/ops/Makefile/k8s/minikube.mk
include $(SELF_DIR)/ops/Makefile/k8s/csi.mk
include $(SELF_DIR)/ops/Makefile/k8s/metallb.mk
include $(SELF_DIR)/ops/Makefile/k8s/prometheus.mk
include $(SELF_DIR)/ops/Makefile/k8s/skaffold.mk
include $(SELF_DIR)/ops/Makefile/k8s/telepresence.mk
