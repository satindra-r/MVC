JWT=$(curl -s -X POST http://localhost:8090/api/user/login \
  -H "Username: admin" \
  -H "Password: password" \
  -H "Content-Length: 0")

ab -k -l -n 100000 -c 1000 -H "Cookie: JWT=$JWT"  http://127.0.0.1:8090/items

ab -n 1000 -c 100  -p /dev/null -T "application/json" -H "Cookie: JWT=$JWT" -H "SectionName: E" http://127.0.0.1:8090/api/sections