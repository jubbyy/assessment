###
Invoke-RestMethod  -uri http://localhost:2565/expense -Method POST -Body '{"title": "post","amount":2656,"note":"notees","tags":["tag1","tag2"]}'

Invoke-RestMethod  -uri http://localhost:2565/expense/1

Invoke-RestMethod  -uri http://localhost:2565/expense/1 -Method PUT -Body '{"title": "post","amount":2657,"note":"notees","tags":["tag1","tag2","tags3"]}'

Invoke-RestMethod  -uri http://localhost:2565/expense/1 -Method DELETE