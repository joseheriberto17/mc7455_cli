herramienta para el diagnostico y configuracion de dispositivo modem sierra MC7455

capacidades

- enviar comando AT por el puerto serial que ofrece mc7455

comando principal para compilar y transferir a equipo objetivo

GOARM=7 GOARCH=arm go build ./cmd/mc7455_cli ; scp -P 2122 ./mc7455_cli root@10.1.162.175:/SD/mc7455_cli
