# iban-validation

### Summary
A service to valid iban number. it verifies the IBAN number with following these steps:
1. the specific length of the IBAN for each country's standard
2. country specific IBAN structure
3. IBAN CHECKSUM

*NOTE
1. Some countries use internal check digit algorithms to validate domestic BBAN, Every country uses a different algorithm and in some countries algorithms vary from bank to bank or even individual branches. that is not included in this service
2. supported country can be found here https://github.com/lilin1013/iban-validation/blob/main/iban/country_setting.go




### How to start
1. go to terminal, and start the service by `make run`
2. go to another terminal window and 
`curl --location --request POST '127.0.0.1:8080/valid' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ibanNumber": "GB33BUKB20201555555555"
}'`
