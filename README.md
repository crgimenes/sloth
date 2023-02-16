# sloth

# debian Postgres

```bash
virt-install \
    --name postgres \
    --os-type linux \
    --os-variant debian9 \
    --network bridge=br0 \
    --ram 1024 \
    --disk /kvm/disk/postgres.img,device=disk,bus=virtio,size=10,format=qcow2 \
    --nographics \
    --hvm \
    --location /kvm/iso/debian-11.3.0-amd64-netinst.iso \
    --extra-args "console=ttyS0"
```

# debian RabbitMQ

```bash
virt-install \
    --name rabbitmq \
    --os-type linux \
    --os-variant debian9 \
    --network bridge=br0 \
    --ram 1024 \
    --disk /kvm/disk/rabbitmq.img,device=disk,bus=virtio,size=10,format=qcow2 \
    --nographics \
    --hvm \
    --location /kvm/iso/debian-11.3.0-amd64-netinst.iso \
    --extra-args "console=ttyS0" 
```

<http://rabbitmq.local:15672/>
