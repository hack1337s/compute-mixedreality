function setgoenv() {
	$workspace = "/Users/alexandrerochaix/Desktop/Compute_Server/compute-mixedreality/WebApp"
	export GOBIN = "$workspace/bin"
	export GOPATH = "$workspace"
	export GOROOT = "$workspace/go"
}
export -f setgoenv
