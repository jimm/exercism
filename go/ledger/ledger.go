package ledger

import (
	"fmt"
	"errors"
	"strings"
	"sort"
)

const testVersion = 3

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type localization struct {
	date, description, change, dateSep string
	negPrefix, negSuffix, afterCurrency, thousands, decimal string
	dateFieldIndexes [3]int
}

var localizations = map[string]localization{
	"en-US": localization{
		date: "Date", description: "Description", change: "Change",
		dateSep: "/", dateFieldIndexes: [3]int{1, 2, 0},
		negPrefix: "(", negSuffix: ")", afterCurrency: "", thousands: ",", decimal: "."},
	"nl-NL": localization{
		date: "Datum", description: "Omschrijving", change: "Verandering",
		dateSep: "-", dateFieldIndexes: [3]int{2, 1, 0},
		negPrefix: "", negSuffix: "-", afterCurrency: " ", thousands: ".", decimal: ","},
}

var currencies = map[string]string{
	"USD": "$",
	"EUR": "€",
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	sortEntries(entriesCopy)

	loc, ok := localizations[locale]
	if !ok {
		return "", errors.New("")
	}
	currencySymbol, ok := currencies[currency]
	if !ok {
		return "", errors.New("")
	}

	ledgerLines := make([]string, len(entries) + 1)
	ledgerLines[0] = fmt.Sprintf("%-10s | %-25s | %s", loc.date, loc.description, loc.change)
	for i, entry := range entriesCopy {
		d, err := formatDate(entry.Date, loc)
		if err != nil {
			return "", errors.New("")
		}
		a := formatMoney(entry.Change, currencySymbol, loc)
		ledgerLines[i+1] = fmt.Sprintf("%-10s | %-25s | %13s", d,
			truncate(entry.Description, 25), a)
	}
	return strings.Join(ledgerLines, "\n") + "\n", nil
}

func formatDate(date string, loc localization) (string, error) {
	if len(date) != 10 {
		return "", errors.New("")
	}
	fields := strings.SplitN(date, "-", 3)
	if len(fields) < 3 {
		return "", errors.New("")
	}
	d := fmt.Sprintf("%s%s%s%s%s", fields[loc.dateFieldIndexes[0]], loc.dateSep,
		fields[loc.dateFieldIndexes[1]], loc.dateSep,
		fields[loc.dateFieldIndexes[2]])
	return d, nil
}

func truncate(s string, maxlen int) string {
	if len(s) <= 25 {
		return s
	}
	return s[:22] + "..."
}

func formatMoney(cents int, currencySymbol string, loc localization) string {
	prefix := ""
	suffix := " "
	if cents < 0 {
		cents = -cents
		prefix = loc.negPrefix
		suffix = loc.negSuffix
	}

	centsStr := fmt.Sprintf("%03d", cents)
	rest := centsStr[:len(centsStr)-2]
	var parts []string
	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}
	parts = append(parts, rest)

	return fmt.Sprintf("%s%s%s%s%s%s%s",
		prefix, currencySymbol, loc.afterCurrency,
		strings.Join(reverse(parts), loc.thousands),
		loc.decimal, centsStr[len(centsStr)-2:], suffix)
}

func reverse(strs []string) []string {
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
strs[i], strs[j] = strs[j], strs[i]
	}
	return strs
}

// **************** sorting ****************

type entrySort []Entry

func (es entrySort) Len() int { return len(es) }

func (es entrySort) Swap(i, j int) { es[i], es[j] = es[j], es[i] }

func (es entrySort) Less(i, j int) bool {
	if es[i].Date < es[j].Date {
		return true
	} else if es[i].Description < es[j].Description {
		return true
	} else if es[i].Change < es[j].Change {
		return true
	}
	return false
}

func sortEntries(entries []Entry) {
	sort.Sort(entrySort(entries))
}
