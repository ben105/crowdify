CREATE KEYSPACE IF NOT EXISTS crowdify WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

CREATE ROLE [replace me] WITH PASSWORD = 'replace me'
    AND SUPERUSER = true 
    AND LOGIN = true;
