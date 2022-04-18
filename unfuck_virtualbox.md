```cmd
bcdedit /set hypervisorlaunchtype off
```

```
systeminfo
```

if `systeminfo` or `msinfo32` shows the verbose Hyper-V entries:
```
Anforderungen für Hyper-V:                     Erweiterungen für den VM-Überwachungsmodus: Ja
                                               Virtualisierung in Firmware aktiviert: Ja
                                               Adressübersetzung der zweiten Ebene: Ja
                                               Datenausführungsverhinderung verfügbar: Ja
```
youre fine :)
