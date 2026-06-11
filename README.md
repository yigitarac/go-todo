# GO-TODO

Go ile yazılmış, aşırı hızlı ve minimal bir Görev Yöneticisi (Todo).

## Kurulum
Sisteminde Go kuruluysa, repoyu indirip doğrudan kendi lokal komut dizinine derlemen yeterli:

    git clone [https://github.com/yigitarac/go-todo.git](https://github.com/yigitarac/go-todo.git)
    cd go-todo
    go build -o ~/.local/bin/todo

*(Not: `~/.local/bin` dizininin PATH'e ekli olduğundan emin ol).*

## Kullanım
Terminalde sadece `todo` yazarak her zaman yardım menüsüne ulaşabilirsin. Temel komutlar şöyle:

* **Görev Ekle:** `todo add "Projeyi Bitir" 3` *(3 günlük süreyle görev ekler)*
* **Listele:** `todo ls` veya `todo list` *(Sadece bekleyen görevleri gösterir)*
* **Arşivi Görüntüle:** `todo ls all` *(Bitenler ve süresi geçenler dahil tam listeyi verir)*
* **Tik At:** `todo check 1` veya `todo done 1` *(1. görevi tamamlandı olarak işaretler)*
* **Kalıcı Sil:** `todo rm 1` veya `todo remove 1` *(1. görevi çöpe atar)*
