just := env_var_or_default("JUST", just_executable())
wash := env_var_or_default("WASH", "wash")
docker := env_var_or_default("DOCKER", "docker")
wit-deps := env_var_or_default("WIT_DEPS", "wit-deps")
devcontainer := env_var_or_default("DEVCONTAINER", "devcontainer")

devcontainer_container_name := env_var_or_default("DEVCONTAINER_CONTAINER_NAME", "kcon2024-workshop")

@_default:
    {{just}} --list

# Ensure `wash` tooling is present
@_ensure-tool-wash:
    command -v {{wash}} || echo "wash binary is missing/not on your PATH, please install it (see: https://wasmcloud.com/docs/installation)"

# Ensure `docker` tooling is present
@_ensure-tool-docker:
    command -v {{docker}} || echo "docker binary is missing/not on your PATH, please install it (see: https://docs.docker.com/get-started/get-docker/)"

# Ensure `devcontainer` tooling is present
@_ensure-tool-devcontainer:
    command -v {{devcontainer}} || echo "devcontainer binary is missing/not on your PATH, please install it (see: https://github.com/devcontainers/cli)"

# Ensure `wit-deps` tooling is present
@_ensure-tool-wit-deps:
    command -v {{wit-deps}} || echo "wit-deps binary is missing/not on your PATH, please install it (see: https://github.com/bytecodealliance/wit-deps)"

###############
# Environment #
###############

# Check for required tools
setup-manual: _ensure-tool-wash _ensure-tool-docker

# Start the local devcontainer
start-devcontainer: _ensure-tool-devcontainer
    {{devcontainer}} up --workspace-folder .

# Stop the local devcontainer
stop-devcontainer: _ensure-tool-docker
    {{docker}} stop {{devcontainer_container_name}}
    {{docker}} rm {{devcontainer_container_name}}

#########
# Build #
#########

# Fetch WIT interfaces
fetch-wit:
    {{wit-deps}} && {{wit-deps}} update

# Build the `nudb` WebAssembly component
build: _ensure-tool-wash fetch-wit
    {{wash}} build -p nudb/wasmcloud.toml
