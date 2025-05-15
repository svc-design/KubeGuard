#!/bin/bash
set -e

VERSION="$1"
ARCH="$2"

[[ -z "$VERSION" || -z "$ARCH" ]] && echo "Usage: $0 <version> <amd64|arm64>" && exit 1

mkdir -p offline_package/bin offline_package/scripts

# 下载二进制文件
wget "https://github.com/<your-github-user>/k8s-backup-tool/releases/download/${VERSION}/k8s-backup-tool-linux-${ARCH}" -O offline_package/bin/k8s-backup-tool
chmod +x offline_package/bin/k8s-backup-tool

# 下载依赖工具（jq、yq、velero）
wget -qO offline_package/bin/yq https://github.com/mikefarah/yq/releases/latest/download/yq_linux_${ARCH}
chmod +x offline_package/bin/yq

wget -qO offline_package/bin/velero.tar.gz https://github.com/vmware-tanzu/velero/releases/latest/download/velero-linux-${ARCH}.tar.gz
tar -zxvf offline_package/bin/velero.tar.gz -C offline_package/bin --strip-components=1 velero-*/velero
rm offline_package/bin/velero.tar.gz
chmod +x offline_package/bin/velero

# 提供安装脚本
cat > offline_package/scripts/install.sh <<EOF
#!/bin/bash
sudo cp bin/* /usr/local/bin/
echo "✅ 安装完成"
EOF

chmod +x offline_package/scripts/install.sh

# 打包离线安装包
tar czf k8s-backup-tool-offline-${ARCH}.tar.gz offline_package

echo "✅ 离线安装包生成：k8s-backup-tool-offline-${ARCH}.tar.gz"

