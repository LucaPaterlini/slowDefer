#  Comparison of the same function with or without the defer


## Benchmarks

### Go 
``` go version go1.12.7 linux/amd64
```

### Hw
```
H/W path       Device     Class          Description
 ====================================================
                           system         Inspiron 7570 (07EA)
 /0                        bus            09KT9M
 /0/0                      memory         64KiB BIOS
 /0/3d                     memory         8GiB System Memory
 /0/3d/0                   memory         8GiB SODIMM DDR4 Synchronous Unbuffered (Unregistered) 2400 MHz (0.4 ns)
 /0/3d/1                   memory         Project-Id-Version: lshwReport-Msgid-Bugs-To: FULL NAME <EMAIL@ADDRESS>PO-Revi
 /0/41                     memory         256KiB L1 cache
 /0/42                     memory         1MiB L2 cache
 /0/43                     memory         8MiB L3 cache
 /0/44                     processor      Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
 /0/100                    bridge         Xeon E3-1200 v6/7th Gen Core Processor Host Bridge/DRAM Registers
 /0/100/2                  display        UHD Graphics 620
 /0/100/4                  generic        Xeon E3-1200 v5/E3-1500 v5/6th Gen Core Processor Thermal Subsystem
 /0/100/14                 bus            Sunrise Point-LP USB 3.0 xHCI Controller
 /0/100/14/0    usb1       bus            xHCI Host Controller
 /0/100/14/0/2             input          USB Receiver
 /0/100/14/0/5             multimedia     Integrated_Webcam_HD
 /0/100/14/0/7             communication  Bluetooth wireless interface
 /0/100/14/1    usb2       bus            xHCI Host Controller
 /0/100/14.2               generic        Sunrise Point-LP Thermal subsystem
 /0/100/15                 generic        Sunrise Point-LP Serial IO I2C Controller #0
 /0/100/15.1               generic        Sunrise Point-LP Serial IO I2C Controller #1
 /0/100/16                 communication  Sunrise Point-LP CSME HECI #1
 /0/100/17                 storage        Sunrise Point-LP SATA Controller [AHCI mode]
 /0/100/1c                 bridge         Sunrise Point-LP PCI Express Root Port #1
 /0/100/1c/0               display        GM108M [GeForce MX130]
 /0/100/1c.4               bridge         Sunrise Point-LP PCI Express Root Port #5
 /0/100/1c.4/0  enp2s0     network        RTL8111/8168/8411 PCI Express Gigabit Ethernet Controller
 /0/100/1c.5               bridge         Sunrise Point-LP PCI Express Root Port #6
 /0/100/1c.5/0  wlp3s0     network        Wireless 7265
 /0/100/1f                 bridge         Intel(R) 100 Series Chipset Family LPC Controller/eSPI Controller - 9D4E
 /0/100/1f.2               memory         Memory controller
 /0/100/1f.3               multimedia     Sunrise Point-LP HD Audio
 /0/100/1f.4               bus            Sunrise Point-LP SMBus
```
### results

```
goos: linux
goarch: amd64
BenchmarkCallNoDefer-8   	2000000000	         0.27 ns/op
BenchmarkCallDefer-8     	30000000	        42.4 ns/op
PASS
```

## Conclusion

write as issue your conclusions
