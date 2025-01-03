!/bin/bash

# Configurações do banco de dados
DB_NAME="seu_banco"           # Nome do banco de dados
DB_USER="seu_usuario"         # Usuário do banco
DB_PASSWORD="sua_senha"       # Senha do banco
DB_HOST="localhost"           # Host do banco
DB_PORT="5432"                # Porta do banco

# Conectar no banco
export PGPASSWORD=$DB_PASSWORD

echo "Criando tabelas no banco de dados $DB_NAME..."

psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME <<EOSQL

-- Criar tabela Borrower se não existir
CREATE TABLE IF NOT EXISTS Borrower (
    borrower_id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(255) UNIQUE,
    cpf VARCHAR(14) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Criar tabela Simulation se não existir
CREATE TABLE IF NOT EXISTS Simulation (
    simulation_id UUID PRIMARY KEY,
    borrower_id UUID REFERENCES Borrower(borrower_id) ON DELETE CASCADE,
    loan_value NUMERIC(15, 2) NOT NULL,
    number_of_installments INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    status VARCHAR(50) NOT NULL,
    interest_rate NUMERIC(5, 2) NOT NULL
);

-- Criar tabela Setup se não existir
CREATE TABLE IF NOT EXISTS Setup (
    setup_id UUID PRIMARY KEY,
    capital NUMERIC(15, 2) NOT NULL,
    fees NUMERIC(5, 2) NOT NULL,
    interest_rate NUMERIC(5, 2) NOT NULL,
    escope VARCHAR(255) NOT NULL,
    escope_is_valid BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

EOSQL

echo "Tabelas criadas (se necessário)!"
