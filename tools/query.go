package tools

const (

	//Services
	GET_ALL_DATA_SERVICES   = "SELECT * FROM services"
	GET_DATA_SERVICES_BY_ID = "SELECT * FROM services WHERE serviceID = ?"
	INSERT_DATA_SERVICES    = "INSERT INTO services VALUES (?, ?, ?, ?)"
	UPDATE_DATA_SERVICES    = `UPDATE services SET serviceType = ?, vehicleType = ?, servicePrice = ?
	 WHERE serviceID = ?`
	DELETE_DATA_SERVICES = "DELETE FROM services WHERE serviceID = ?"

	//customers
	GET_ALL_DATA_CUSTOMERS   = "SELECT * FROM customers"
	GET_DATA_CUSTOMERS_BY_ID = "SELECT * FROM customers WHERE customerID = ?"
	INSERT_DATA_CUSTOMERS    = "INSERT INTO customers VALUES (?, ?, ?, ?, ?, ?)"
	UPDATE_DATA_CUSTOMERS    = `UPDATE customers SET customerName = ?, bookingDate = ?, bookingTime = ?, 
	numberPlat = ?, serviceID = ? WHERE customerID = ?`
	DELETE_DATA_CUSTOMERS = "DELETE FROM customers WHERE customerID = ?"

	//progress
	GET_ALL_PROGRESS        = "SELECT * FROM progress"
	GET_DATA_PROGRESS_BY_ID = "SELECT * FROM progress WHERE progressID = ?"
	INSERT_DATA_PROGRESS    = "INSERT INTO progress VALUES (?, ?, ?, ?)"
	UPDATE_DATA_PROGRESS    = `UPDATE progress SET serviceID = ?, customerID = ?, serviceStatus = ?
	 WHERE progressID = ?`
	DELETE_DATA_PROGRESS = "DELETE FROM progress WHERE progressID = ?"

	// GET_TOTAL_SALARY_PERMONTH = `select(select sum(kartap.gaji_pokok + kartap.tunjangan) from kartap)
	// + (select sum(karkontrak.gaji_pokok) from karkontrak) as "Total Gaji PerBulan"`
	// GET_TOTAL_SALARY_PEREMPLOYEE = ` select kartap.nama_lengkap, (kartap.gaji_pokok + kartap.tunjangan) as "total gaji" from kartap
	// union select karkontrak.nama_lengkap, (karkontrak.gaji_pokok) "total gaji" from karkontrak`
	// GET_TOTAL_FIX_EMPLOYEE       = `select count(*)as "total karyawan tetap"from kartap`
	// GET_TOTAL_TEMPORARY_EMPLOYEE = `select count(*)as "total karyawan kontrak"from karkontrak`
	// GET_TOTAL_ACTIVE_EMPLOYEE    = `select (select count(*)from kartap where status_karyawan="aktif") + (select count(*)from karkontrak where status_karyawan = "aktif") as "total karyawan aktif"`
)
