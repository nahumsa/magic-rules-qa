<h1 align="center">Magic: the Gathering Rules Q&A</h1>

Welcome to the **Q&A with Magic: The Gathering Rules** project! This tool is designed to help Magic: The Gathering (MTG) players quickly find accurate answers to rules-related questions.
Whether you're a beginner or a seasoned player, this project aims to make navigating the complex rules of MTG easier and more intuitive.

## Features

### Completed

- **Comprehensive Rules Database:** Access an extensive collection of rules and official rulings directly from the MTG Comprehensive Rules.
- **Search Functionality:** Use keywords or phrases to find relevant rules and explanations.
- **Vector Search:** Utilize vector search technology to quickly retrieve and rank the most relevant information from the rules database.
- **Regular Updates:** Stay up to date with the latest rules changes and errata from Wizards of the Coast.

### In progress

- **Instant Q&A:** Get immediate answers to specific questions by querying the rules database.
- **GenAI Integration:** Leverage GenAI for advanced natural language understanding to enhance the accuracy and relevance of responses.
- **User-Friendly Interface:** Clean and intuitive design for easy navigation and interaction.

## Getting Started

### Prerequisites

To run this project, you'll need:

- [Golang 1.22+](https://golang.org/dl/)
- [jina.ai](https://jina.ai/embeddings) (for embeddings)
- [QDRANT](https://qdrant.tech/) (for efficient information retrieval)

## Environment Variables

| Variable Name  | Description                                                  | Example Value               |
|----------------|--------------------------------------------------------------|-----------------------------|
| `JINA_API_KEY` | API key for jina.ai                                          | `jina_naiasokdlnmasnd123ln` |
