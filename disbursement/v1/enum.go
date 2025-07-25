// Copyright 2025 PT Espay Debit Indonesia Koe
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package disbursement

// Code generated by generate-enum.sh; DO NOT EDIT.

type Latesttransactionstatus string
const (
    LATESTTRANSACTIONSTATUS_00_ Latesttransactionstatus = "00"
    LATESTTRANSACTIONSTATUS_01_ Latesttransactionstatus = "01"
    LATESTTRANSACTIONSTATUS_02_ Latesttransactionstatus = "02"
    LATESTTRANSACTIONSTATUS_03_ Latesttransactionstatus = "03"
    LATESTTRANSACTIONSTATUS_04_ Latesttransactionstatus = "04"
    LATESTTRANSACTIONSTATUS_05_ Latesttransactionstatus = "05"
    LATESTTRANSACTIONSTATUS_06_ Latesttransactionstatus = "06"
    LATESTTRANSACTIONSTATUS_07_ Latesttransactionstatus = "07"
)

type Sourceplatform string
const (
    SOURCEPLATFORM_IPG_ Sourceplatform = "IPG"
)

type Terminaltype string
const (
    TERMINALTYPE_APP_ Terminaltype = "APP"
    TERMINALTYPE_WEB_ Terminaltype = "WEB"
    TERMINALTYPE_WAP_ Terminaltype = "WAP"
    TERMINALTYPE_SYSTEM_ Terminaltype = "SYSTEM"
)

type Orderterminaltype string
const (
    ORDERTERMINALTYPE_APP_ Orderterminaltype = "APP"
    ORDERTERMINALTYPE_WEB_ Orderterminaltype = "WEB"
    ORDERTERMINALTYPE_WAP_ Orderterminaltype = "WAP"
    ORDERTERMINALTYPE_SYSTEM_ Orderterminaltype = "SYSTEM"
)

type Paymethod string
const (
    PAYMETHOD_BALANCE_ Paymethod = "BALANCE"
    PAYMETHOD_COUPON_ Paymethod = "COUPON"
    PAYMETHOD_NET_BANKING_ Paymethod = "NET_BANKING"
    PAYMETHOD_CREDIT_CARD_ Paymethod = "CREDIT_CARD"
    PAYMETHOD_DEBIT_CARD_ Paymethod = "DEBIT_CARD"
    PAYMETHOD_VIRTUAL_ACCOUNT_ Paymethod = "VIRTUAL_ACCOUNT"
    PAYMETHOD_OTC_ Paymethod = "OTC"
    PAYMETHOD_DIRECT_DEBIT_CREDIT_CARD_ Paymethod = "DIRECT_DEBIT_CREDIT_CARD"
    PAYMETHOD_DIRECT_DEBIT_DEBIT_CARD_ Paymethod = "DIRECT_DEBIT_DEBIT_CARD"
    PAYMETHOD_ONLINE_CREDIT_ Paymethod = "ONLINE_CREDIT"
    PAYMETHOD_LOAN_CREDIT_ Paymethod = "LOAN_CREDIT"
    PAYMETHOD_NETWORK_PAY_ Paymethod = "NETWORK_PAY"
)

type Payoption string
const (
    PAYOPTION_NETWORK_PAY_PG_SPAY_ Payoption = "NETWORK_PAY_PG_SPAY"
    PAYOPTION_NETWORK_PAY_PG_OVO_ Payoption = "NETWORK_PAY_PG_OVO"
    PAYOPTION_NETWORK_PAY_PG_GOPAY_ Payoption = "NETWORK_PAY_PG_GOPAY"
    PAYOPTION_NETWORK_PAY_PG_LINKAJA_ Payoption = "NETWORK_PAY_PG_LINKAJA"
    PAYOPTION_NETWORK_PAY_PG_CARD_ Payoption = "NETWORK_PAY_PG_CARD"
    PAYOPTION_VIRTUAL_ACCOUNT_BCA_ Payoption = "VIRTUAL_ACCOUNT_BCA"
    PAYOPTION_VIRTUAL_ACCOUNT_BNI_ Payoption = "VIRTUAL_ACCOUNT_BNI"
    PAYOPTION_VIRTUAL_ACCOUNT_MANDIRI_ Payoption = "VIRTUAL_ACCOUNT_MANDIRI"
    PAYOPTION_VIRTUAL_ACCOUNT_BRI_ Payoption = "VIRTUAL_ACCOUNT_BRI"
    PAYOPTION_VIRTUAL_ACCOUNT_BTPN_ Payoption = "VIRTUAL_ACCOUNT_BTPN"
    PAYOPTION_VIRTUAL_ACCOUNT_CIMB_ Payoption = "VIRTUAL_ACCOUNT_CIMB"
    PAYOPTION_VIRTUAL_ACCOUNT_PERMATA_ Payoption = "VIRTUAL_ACCOUNT_PERMATA"
)

type Type string
const (
    TYPE_PAY_RETURN_ Type = "PAY_RETURN"
    TYPE_NOTIFICATION_ Type = "NOTIFICATION"
)

type Actortype string
const (
    ACTORTYPE_USER_ Actortype = "USER"
    ACTORTYPE_MERCHANT_ Actortype = "MERCHANT"
    ACTORTYPE_MERCHANT_OPERATOR_ Actortype = "MERCHANT_OPERATOR"
    ACTORTYPE_BACK_OFFICE_ Actortype = "BACK_OFFICE"
    ACTORTYPE_SYSTEM_ Actortype = "SYSTEM"
)

