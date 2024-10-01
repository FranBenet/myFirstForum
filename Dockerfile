from golang:1.21
workdir /
copy forum.db literary-lions /
copy dbaser dbaser/
copy handlers handlers/
copy helpers helpers/
copy middleware middleware/
copy models models/
copy session session/
copy web web/
expose 8080
cmd ./literary-lions
