!/bin/bash

# Configurações do banco de dados
# deveriam estar em um cofre de senha para a segurança, os dados nao podem ser expostos, porem como e so um case tecnico eu deixei assim
# caso tenha algum dado de cartao na tabela ou conta estes dados sigilosos  devem ser criptografados
#  sempre garantido a segurança !!
DB_NAME="simulations"           # Nome do banco de dados
DB_USER="postgres"         # Usuário do banco
DB_PASSWORD="12345"       # Senha do banco
DB_HOST="localhost"           # Host do banco
DB_PORT="5432"                # Porta do banco

# Conectar no banco
export PGPASSWORD=$DB_PASSWORD

echo "Criando tabelas no banco de dados $DB_NAME..."

psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME <<EOSQL

-- Criar tabela Borrower se não existir
CREATE TABLE IF NOT EXISTS Borrower  (
    borrower_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(255) UNIQUE,
    cpf VARCHAR(14) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Criar tabela Simulation se não existir
CREATE TABLE IF NOT EXISTS Simulation (
    simulation_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    borrower_id UUID REFERENCES Borrower(borrower_id) ON DELETE CASCADE,
    loan_value NUMERIC NOT NULL,
    number_of_installments INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    status VARCHAR(50) NOT NULL,
    interest_rate NUMERIC NOT NULL
);

-- Criar tabela Setup se não existir
CREATE TABLE IF NOT EXISTS Setup (
    setup_id VARCHAR PRIMARY KEY,
    capital NUMERIC NOT NULL,
    fees NUMERIC NOT NULL,
    interest_rate NUMERIC NOT NULL,
    escope VARCHAR(255) NOT NULL,
    escope_is_valid BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

EOSQL

echo "Tabelas criadas (se necessário)!"
