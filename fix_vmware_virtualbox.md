wsl and vmware / virtualbox can't coexist on the same windows system.  
this is a fix to make the hypervisors vmware and virtualbox work again on a windows system thats been absued by wsl or docker.  
the fix is basically setting the `hypervisortype` of that system to `off`, which disables `Hyper-V`.  
if you want to use `wsl` or `docker` on your windows system you havbe to set that value to `auto`.  

```cmd
bcdedit /set hypervisorlaunchtype off
```

```cmd
systeminfo
```

if `systeminfo` or `msinfo32` shows the verbose `Hyper-V` entries:
```log
Anforderungen für Hyper-V:                     Erweiterungen für den VM-Überwachungsmodus: Ja
                                               Virtualisierung in Firmware aktiviert: Ja
                                               Adressübersetzung der zweiten Ebene: Ja
                                               Datenausführungsverhinderung verfügbar: Ja
```

source:  
https://github.com/MicrosoftDocs/WSL/issues/1820
