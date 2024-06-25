# Bewerberchallenge

Hey du! 👋

Willkommen zur TEQWERK Bewerberchallenge! In diesem Projekt hast du die Möglichkeit, dein Können in den Bereichen Infrastruktur, Deployment (CI/CD) und Cloud-Lösungen unter Beweis zu stellen. Dein Ziel ist es, eine hochkritische Krankenhausanwendung zu implementieren, die rund um die Uhr verfügbar sein muss. Hier findest du alle notwendigen Informationen und Aufgaben, um erfolgreich durch die Challenge zu kommen. Falls du gar nicht mehr weiter kommst oder Fragen hast, sind wir natürlich immer für dich da! 

# Inhaltsverzeichnis

- [Bewerberchallenge](#bewerberchallenge)
    - [Szenario](#szenario)
        - [Anforderungen](#anforderungen)
    - [Deine Aufgaben im Projekt](#deine-aufgaben-im-projekt)
        1. [Fork das Repository und richte die Basisanwendung ein](#1-fork-das-repository-und-richte-die-basisanwendung-ein)
        2. [Erstelle eine GitHub Action Terraform Pipeline für das Infrastruktur Deployment](#2-erstelle-eine-github-action-terraform-pipeline-für-das-infrastruktur-deployment)
        3. [Lade die CSV Datei 'patient_data.csv' mit den Daten eines Vorsystems per Data Factory in die SQL Datenbank](#3-lade-die-csv-datei-patient_datacsv-mit-den-daten-eines-vorsystems-per-data-factory-in-die-sql-datenbank)
        4. [Verwende Private Endpoints & Private DNS Zones](#4-verwende-private-endpoints--private-dns-zones)
        5. [Wähle eine geeignete, skalierbare Infrastruktur auf Azure](#5-wähle-eine-geeignete-skalierbare-infrastruktur-auf-azure)
        6. [Konzipiere ein Monitoring für die Applikation sowie Infrastrukur](#6-konzipiere-ein-monitoring-für-die-applikation-sowie-infrastrukur)
- [Bonus Aufgaben](#bonus-aufgaben)
    1. [Entwickle eine Authentifizierung für die Anwendung mit Entra ID oder GitHub App Registration](#1-entwickle-eine-authentifizierung-für-die-anwendung-mit-entra-id-oder-github-app-registration)
    2. [Integration eines User & Access Management Systems](#2-integration-eines-user--access-management-systems)
- [Lokales Deployment](#lokales-deployment)
- [Cloud Deployment mit Terraform](#cloud-deployment-mit-terraform)
    1. [Richte die Infrastruktur mit Terraform ein](#1-richte-die-infrastruktur-mit-terraform-ein)
    2. [Verbinde deine CI/CD Pipeline mit deinem GitHub-Account](#2-verbinde-deine-cicd-pipeline-mit-deinem-github-account)
- [Präsentation und Abgabe](#präsentation-und-abgabe)

## Szenario
Stell dir vor, du arbeitest an der Entwicklung unseres neuen Produkts **TEQWERK Hospital Patients Manager** mit. Diese Anwendung wird von einem großen Krankenhaus genutzt, um Patientendaten zu verwalten. Aufgrund der kritischen Natur der Daten sowie der überlebenswichtig schnellen Abläufe im Krankenhaus, muss die Anwendung äußerst zuverlässig, skalierbar und sicher sein. Bedenke das bei deinen Infrastrukturentscheidungen immer mit!

### Technische Anforderungen
Die Anwendung sollte jederzeit verfügbar sein und soll auf einer skalierbaren Architektur basieren - der Kunde erwartet, dass die Infrastruktur nach einer Migration für ein zukünftiges Wachstum sowie weitere Krankenhäuser vorbereitet ist. Außerdem ist wichtig, dass du beim Deployment auf eine nachvollziehbare Namensgebung der Komponenten achtest und Best Practices aus dem Microsoft Cloud-Adoption-Framework und/oder dem Well-Architected Framework berücksichtigst. Behalte auch Sicherheitsaspekte im Auge. Achte darauf eine geeignete Cloud-Region für das Projekt zu wählen.

## Deine Aufgaben im Projekt

1. **Forke das Repository und richte die Basisanwendung ein:**
        - Starte mit dem Forken des Repositories.
        - Die Anwendung befindet sich im `app` Ordner und kann lokal durch `docker-compose` gestartet werden.
2. **Erstelle eine GitHub Action Terraform Pipeline für das Infrastruktur Deployment:**
        - Implementiere eine CI/CD Pipeline, die den Deployment-Prozess automatisiert.
        - Stelle sicher, dass alle Ressourcen gemäß den Namenskonventionen des Microsoft Cloud Adoption Framework und/oder Well-Architected Framework benannt werden.
        - Die Docker-Container sollten in einer geeigneten Pipeline gebaut und in der GitHub Registry abgelegt werden.
3. **Lade die CSV Datei 'patient_data.csv' mit den Daten eines Vorsystems per Azure Data Factory in die SQL Datenbank:**
        - Implementiere eine Azure Data Factory Pipeline, die die Daten automatisiert in die SQL Datenbank lädt. Verifiziere den erfolgreichen Upload in der Anwendung.
4. **Verwende Private Networking:**
        - Stelle sicher, dass die Datenbank vom Backend privat geroutet wird.
5. **Wähle eine geeignete, skalierbare Infrastruktur auf Azure:**
        - Du bist der Architekt dieses Projekts. Überlege dir, auf welcher Infrastruktur die Anwendung am sinnvollsten betrieben werden sollte. Integriere die Komponenten in dein Architekturbild.
6. **Achte auf die Sicherheit der Anwendung und Kommunikation zwischen den Komponenten:**
        - Arbeite mit Managed Identites um die Authentifierzierung zwischen den Teilen der Applikation zu ermöglichen.
7. **Konzipiere ein Monitoring für die Applikation sowie Infrastrukur**
     - Welche Möglichkeiten gibt es die Anwendungen zu monitoren? Welche Möglichkeiten eignen sich für die Anwendung?
     - Implementiere ein einfaches Monitoring mit Healthchecks auf die Anwendungskomponenten (falls dafür Logging nötig ist, setze dies ebenfalls um)

## Bonus Aufgaben

1. **Entwickle eine Authentifizierung für die Anwendung mit Entra ID oder GitHub App Registration:**
    - Implementiere eine sichere Login-Methode für die Benutzer der Anwendung.
2. **Integration eines User & Access Management Systems:**
    - Baue ein Admin Center, in dem Benutzer und deren Zugriffsrechte verwaltet werden können.

## Lokales Deployment

Zum Testen der Anwendung lokal, führe folgende Schritte aus:

```bash
cd app
docker-compose up --build
```
⚠️ Hinweis: Das Projekt wurde ursprünglich für ARM64 entwickelt. Passe ggf. die Dateien für eine X64 Prozessorarchitektur an.

## Cloud Deployment mit Terraform

1. **Richte die Infrastruktur mit Terraform ein:**
    - Erstelle ein Terraform Repository, um die gesamte Infrastruktur auf Azure bereitzustellen.
    - Achte dabei auf die Einhaltung der Namenskonventionen des Microsoft Cloud Adoption Framework und/oder Well-Architected Framework.
    - Nutze Best-Practices um dein Terraform Repository bestmöglich wartbar zu machen.

2. **Verbinde deine CI/CD Pipeline mit deinem GitHub-Account:**
    - Automatisiere das Deployment, sodass Infrastrukturänderungen per Pipeline ausgerollt werden.

## Präsentation und Abgabe

Nach Abschluss der Aufgaben, präsentiere deine Lösung in einer 30-45 minütigen Pitch-Session. Erstelle eine PDF-Präsentation, in der du deine Architektur- und Serviceentscheidungen erläuterst. Erstelle außerdem ein Architekturdiagramm, idealerweise in [draw.io](http://draw.io/), und stelle deine Lösung als GitHub-Repository bereit.

---

Viel Spaß bei der Challenge! Bei Fragen, melde dich gerne bei uns! 🧑‍💻

Dein TEQWERK-Team 🧡