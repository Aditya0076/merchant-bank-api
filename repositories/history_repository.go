package repositories

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "merchant-bank-api/models"
)

var historyFile = "data/history.json"

// Simpan history ke file JSON
func SaveHistory(log models.History) error {
    logs, err := readHistory()
    if err != nil {
        return err
    }

    logs = append(logs, log)
    return writeHistory(logs)
}

func readHistory() ([]models.History, error) {
    file, err := os.Open(historyFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var history []models.History
    byteValue, _ := ioutil.ReadAll(file)
    json.Unmarshal(byteValue, &history)
    return history, nil
}

func writeHistory(history []models.History) error {
    data, err := json.Marshal(history)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(historyFile, data, 0644)
}
