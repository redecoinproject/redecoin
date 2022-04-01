// Copyright (c) 2011-2014 The Bitcoin Core developers
// Copyright (c) 2017-2019 The Raven Core developers
// Copyright (c) 2020-2021 The redecoin Core developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

#ifndef REDECOIN_QT_REDECOINADDRESSVALIDATOR_H
#define REDECOIN_QT_REDECOINADDRESSVALIDATOR_H

#include <QValidator>

/** Base58 entry widget validator, checks for valid characters and
 * removes some whitespace.
 */
class redecoinAddressEntryValidator : public QValidator
{
    Q_OBJECT

public:
    explicit redecoinAddressEntryValidator(QObject *parent);

    State validate(QString &input, int &pos) const;
};

/** redecoin address widget validator, checks for a valid redecoin address.
 */
class redecoinAddressCheckValidator : public QValidator
{
    Q_OBJECT

public:
    explicit redecoinAddressCheckValidator(QObject *parent);

    State validate(QString &input, int &pos) const;
};

#endif // REDECOIN_QT_REDECOINADDRESSVALIDATOR_H
