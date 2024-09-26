package dbquery

const (
	InsertToUsers  = `INSERT INTO users (username, password, phone, type) VALUES ($1, $2, $3, $4)`
	InsertToOrders = `INSERT INTO orders (
                    hardware,
                    type_failure,
                    description,
                    client_id,
                    executor_id,
                    status
                    ) VALUES ($1, $2, $3, $4, $5, $6)`
	SelectAll            = "SELECT * FROM %s"
	SelectID             = "SELECT * FROM %s WHERE id=$1"
	SelectKey            = "SELECT * FROM %s WHERE %s=$1"
	UpdateUser           = `UPDATE users SET username=$1, phone=$2, type=$3 WHERE id=$4`
	UpdateOrder          = `UPDATE orders SET hardware=$1, updated_at=current_timestamp, type_failure=$2, description=$3, executor_id=$4, status=$5 WHERE id=$6`
	Delete               = "DELETE FROM %s WHERE id=$1"
	CountTypesFailures   = `SELECT type_failure, COUNT(*) AS count FROM orders GROUP BY type_failure`
	CountCompletedOrders = `SELECT count(*) AS count FROM orders where status='done'`
	AverageTime          = `SELECT AVG((strftime('%s', updated_at) - strftime('%s', created_at)) / 3600) AS average_hours FROM orders WHERE status='done'`
)

const CreateUserTable = `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            username VARCHAR(255) NOT NULL,
            password VARCHAR(255) NOT NULL,
            phone VARCHAR(12) NOT NULL,
            tg_chat_id INTEGER UNIQUE,
            type VARCHAR(255) NOT NULL
        )
    `

const CreateOrderTable = `CREATE TABLE IF NOT EXISTS orders (
           id INTEGER PRIMARY KEY,
           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
           hardware VARCHAR(255) NOT NULL,
           type_failure VARCHAR(255) NOT NULL,
           description VARCHAR(255) NOT NULL,
           client_id INTEGER NOT NULL,
           executor_id INTEGER NOT NULL,
           status VARCHAR(255) NOT NULL,
           FOREIGN KEY (client_id) REFERENCES users (chat_id),
           FOREIGN KEY (executor_id) REFERENCES users (chat_id)
        )
    `
