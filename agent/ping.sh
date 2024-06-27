current_pod_name=$(hostname)
current_node_name=$(kubectl get pod $current_pod_name -o jsonpath='{.spec.nodeName}')
nodes=($(kubectl get nodes -o jsonpath='{range .items[*]}{.status.addresses[?(@.type=="InternalIP")].address}{"\n"}{end}'))
nodes_name=($(kubectl get nodes -o jsonpath='{range .items[*]}{.metadata.name}{"\n"}{end}'))
while true; do
    for index in "${!nodes[@]}"; do
        if [ "${nodes_name[index]}" != "$current_node_name" ]; then
            latency=$(ping -c 1 ${nodes[index]} | grep 'time=' | awk -F'time=' '{print $2}' | awk '{print $1}')
            echo "delay{from=\"$current_node_name\",to=\"${nodes_name[index]}\"}" $latency
        fi
    done
    sleep $TIME
done