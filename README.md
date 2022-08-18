# elemental

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/cbd8761ba935460885a1ff9b49497621)](https://www.codacy.com/gh/PaloAltoNetworks/elemental/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=PaloAltoNetworks/elemental&amp;utm_campaign=Badge_Grade) [![Codacy Badge](https://app.codacy.com/project/badge/Coverage/cbd8761ba935460885a1ff9b49497621)](https://www.codacy.com/gh/PaloAltoNetworks/elemental/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=PaloAltoNetworks/elemental&amp;utm_campaign=Badge_Coverage)

> Note: This is a work in progress

Elemental contains the foundations and code generator to translate and use
Regolithe specifications in a Bahamut/Manipulate environment. It also provides
basic definition of a generic request and response and other low level functions
like optimized encoder and decoder, validations and more.

The generated model from Regolithe specifications (when using elegen from
cmd/elegen) will implement the interfaces described in this package to serve as a base
for being used in a bahamut microservice.
