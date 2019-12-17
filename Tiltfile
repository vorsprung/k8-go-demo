print('Hello Tiltfile')
k8s_yaml('stock.yml')
docker_build('vorsprung/stock','.')
k8s_resource('stockapp', port_forwards='9000')