#!/bin/bash

# Read the VERSION argument
VERSION=$1

# Colors and formatting
BOLD="\033[1m"
CYAN="\033[0;36m"
GREEN="\033[0;32m"
RESET="\033[0m"

# Determine the correct home directory and terraform plugin path based on the OS
if [[ "$OSTYPE" == "linux-gnu"* || "$OSTYPE" == "darwin"* ]]; then
  # Linux or macOS
  TERRAFORM_HOME="$HOME/.terraform.d/plugins"
  TERRAFORM_RC="$HOME/.terraformrc"
elif [[ "$OSTYPE" == "cygwin" || "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
  # Windows
  TERRAFORM_HOME="$APPDATA/terraform.d/plugins"
  TERRAFORM_RC="$APPDATA/terraform.rc"
else
  echo -e "${RED}Unsupported OS type: $OSTYPE${RESET}"
  exit 1
fi

# Copy the native KMS crypto shared library alongside the provider binary.
# The keyprotect-go-client SDK looks for the library next to the executable
# at runtime (via getLibraryPath).  Without it the provider process crashes
# when any ibm_kms_cryptounits resource is applied.
OS_ARCH="$(go env GOOS)-$(go env GOARCH)"
KP_LIB_DIR="$(go env GOPATH)/pkg/mod/github.com/!i!b!m/keyprotect-go-client@v0.17.2/dedicated/internal/lib/${OS_ARCH}"
PLUGIN_DIR="${TERRAFORM_HOME}/registry.terraform.io/ibm-cloud/ibm/${VERSION}/$(go env GOOS)_$(go env GOARCH)"

if [[ -d "${KP_LIB_DIR}" ]]; then
  echo -e "\n${GREEN}==> Copying KMS native library to plugin directory...${RESET}"
  cp "${KP_LIB_DIR}"/ibmkmscrypto* "${PLUGIN_DIR}"/
else
  echo -e "\n${CYAN}[WARN] KMS native library not found at ${KP_LIB_DIR}; ibm_kms_cryptounits will not work${RESET}"
fi

# Update the terraform.rc or .terraformrc file using echo
echo -e "\n${GREEN}==> Updating $TERRAFORM_RC file...${RESET}"
cat <<EOL > "$TERRAFORM_RC"
provider_installation {
    filesystem_mirror {
        path    = "$TERRAFORM_HOME"
        include = ["ibm-cloud/ibm"]
    }
    direct {
        include = ["*/*"]
    }
}
EOL

# Print the instructions for updating the Terraform example
echo -e "\n${CYAN}${BOLD}⬇ Add the following to your Terraform configuration:${RESET}\n"
echo -e "terraform {"
echo -e "  required_providers {"
echo -e "    ibm = {"
echo -e "      source  = \"${CYAN}IBM-Cloud/ibm${RESET}\""
echo -e "      version = \"${CYAN}$VERSION${RESET}\""
echo -e "    }"
echo -e "  }"
echo -e "}\n"
