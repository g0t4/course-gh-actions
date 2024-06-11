namespace tests;

public class Tests
{
    [Test]
    public void Test_10C_Is50F()
    {
        var forecast = new WeatherForecast(DateOnly.FromDateTime(DateTime.Now), 10, "Freezing");
        Assert.That(forecast.TemperatureC, Is.EqualTo(10));
        Assert.That(forecast.TemperatureF, Is.EqualTo(50));
    }
}
