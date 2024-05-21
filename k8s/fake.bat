@echo off
for /L %%i in (1,1,1000) do (
curl --header "Connection: keep-alive" "http://34.65.129.115:8080/"
)
