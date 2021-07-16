# Rebuilding Git With Golang

## Bab 1

### Apa itu git ?

Git adalah suatu tool yang sering kali digunakan untuk pengembangan software. Fungsinya adalah sebagai sistem pengontrol versi (*Version Control System*) pada proyek perangkat lunak.



### Kenapa kita karus menggunakan git?







## Bab 2

### Perkenalan directory .git

Langkah pertama saat kita menggunakan git biasanya menggunakan comand `git init`, comand tersebut akan membuat folder baru dengan nama `.git`.  untuk struktur foldernya akan jadi seperti ini :

``` bash
.git
|- config
|- description
|- HEAD
|- hooks
|  |- update.sample
|- info
|  |- exclude
|- objects
|  |- info
|  |- pack
|- refs
   |- heads
   |- tags
```

