pushd cmd/manager/ >/dev/null
export WATCH_NAMESPACE=default
dlv debug --headless --listen=:2345 --api-version=2
popd >/dev/null