public class Factory {
    public getEmployee(String employeeType) {
        if("commission".equals(employeeType)) {
            return new CommissionedEmployee();
        }
        if("hour".equals(employeeType)) {
            return new HourlyEmployee();
        }
        if("salary".equals(employeeType)) {
            return new SalariedEmployee();
        }
    }
}