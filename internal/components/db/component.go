package db

type Connection struct {
	connectionString string
	logger           logger
}

func New(config config, logger logger) *Connection {
	return &Connection{
		connectionString: config.GetConnectionString(),
		logger:           logger,
	}
}

func (c *Connection) Select() {
	c.logger.Log("select " + c.connectionString)
}

func (c *Connection) Insert() {
	c.logger.Log("insert " + c.connectionString)
}

func (c *Connection) Update() {
	c.logger.Log("update " + c.connectionString)
}

func (c *Connection) Delete() {
	c.logger.Log("delete " + c.connectionString)
}
