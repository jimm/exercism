package ledger

import (
	"fmt"
	"errors"
	"strconv"
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
	negPrefix, negSuffix, thousands, decimal string
	dateFieldIndexes [3]int
}

var localizations = map[string]localization{
	"en-US": localization{date: "Date", description: "Description",
		change: "Change", dateSep: "/", dateFieldIndexes: [3]int{1, 2, 0},
		negPrefix: "(", negSuffix: ")", thousands: ",", decimal: "."},
	"nl-NL": localization{date: "Datum", description: "Omschrijving",
		change: "Verandering", dateSep: "-", dateFieldIndexes: [3]int{2, 1, 0},
		negPrefix: "", negSuffix: "-", thousands: ".", decimal: ","},
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

	s := fmt.Sprintf("%-10s | %-25s | %s\n", loc.date, loc.description, loc.change)

	for _, entry := range entriesCopy {
		d, err := formatDate(entry.Date, loc)
		if err != nil {
			return "", errors.New("")
		}

		negative := false
		cents := entry.Change
		if cents < 0 {
			cents = cents * -1
			negative = true
		}
		var a string
		if locale == "nl-NL" {
			a += currencySymbol
			a += " "
			centsStr := strconv.Itoa(cents)
			switch len(centsStr) {
			case 1:
				centsStr = "00" + centsStr
			case 2:
				centsStr = "0" + centsStr
			}
			rest := centsStr[:len(centsStr)-2]
			var parts []string
			for len(rest) > 3 {
				parts = append(parts, rest[len(rest)-3:])
				rest = rest[:len(rest)-3]
			}
			if len(rest) > 0 {
				parts = append(parts, rest)
			}
			for i := len(parts) - 1; i >= 0; i-- {
				a += parts[i] + "."
			}
			a = a[:len(a)-1]
			a += ","
			a += centsStr[len(centsStr)-2:]
			if negative {
				a += "-"
			} else {
				a += " "
			}
		} else if locale == "en-US" {
			if negative {
				a += "("
			}
			a += currencySymbol
			centsStr := strconv.Itoa(cents)
			switch len(centsStr) {
			case 1:
				centsStr = "00" + centsStr
			case 2:
				centsStr = "0" + centsStr
			}
			rest := centsStr[:len(centsStr)-2]
			var parts []string
			for len(rest) > 3 {
				parts = append(parts, rest[len(rest)-3:])
				rest = rest[:len(rest)-3]
			}
			if len(rest) > 0 {
				parts = append(parts, rest)
			}
			for i := len(parts) - 1; i >= 0; i-- {
				a += parts[i] + ","
			}
			a = a[:len(a)-1]
			a += "."
			a += centsStr[len(centsStr)-2:]
			if negative {
				a += ")"
			} else {
				a += " "
			}
		} else {
			return "", errors.New("")
		}
		var al int
		for _ = range a {
			al++
		}
		s += fmt.Sprintf("%-10s | %-25s | %13s\n", d, truncate(entry.Description, 25), a)
	}
	return s, nil
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
