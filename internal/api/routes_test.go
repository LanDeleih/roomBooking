package api

import (
	"bytes"
	"encoding/json"
	"github.com/lanDeleih/roomBooking/internal/roomRepository"
	"net/http"
	"testing"
	"time"
)

func TestCreateRoom(t *testing.T) {
	c := &http.Client{Timeout: 5 *time.Second}

	body := &roomRepository.Room{
		Number: 1,
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/room/create", buf)
	if err != nil {
		t.Fatalf("req %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf("resp %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %s", err)
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("body %s", err)
	}
}
func TestCreateReservation(t *testing.T) {
	c := &http.Client{Timeout: 5 *time.Second}

	body := &roomRepository.Appointment{
		Customer:  "John",
		StartDate: "10/24/2020 03:16 PM",
		Duration:  "1h",
		Room:      1,
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/booking/create", buf)
	if err != nil {
		t.Fatalf("req %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf("resp %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %s", err)
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("body %s", err)
	}
}
func TestDuplicateReservation(t *testing.T) {
	c := &http.Client{Timeout: 5 *time.Second}

	body := &roomRepository.Appointment{
		Customer:  "John",
		StartDate: "10/24/2020 04:13 PM",
		Duration:  "1h",
		Room:      1,
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/booking/create", buf)
	if err != nil {
		t.Fatalf("req %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf("resp %s", err)
	}

	if resp.StatusCode == http.StatusOK {
		t.Fatalf("status %s", err)
	}
	err = resp.Body.Close()
	if err != nil {
		t.Fatalf("body %s", err)
	}
}
func TestDeleteRoom(t *testing.T) {
	c := &http.Client{Timeout: 5 *time.Second}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/room/delete/1", nil)
	if err != nil {
		t.Fatalf("req %s", err)
	}
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf("resp %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %s", err)
	}
}
func TestGetReservation(t *testing.T) {
	c := &http.Client{Timeout: 5 *time.Second}

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/booking/", nil)
	if err != nil {
		t.Fatalf("req %s", err)
	}
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf("resp %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %s", err)
	}
}
func TestGetRooms(t *testing.T) {
	c := &http.Client{Timeout: 5 *time.Second}

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/room/", nil)
	if err != nil {
		t.Fatalf("req %s", err)
	}
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf("resp %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %s", err)
	}
}
