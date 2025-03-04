import { useState, useEffect } from 'react';

export default function WeatherDisplay({ weatherData, error }) {
  const [currentTime, setCurrentTime] = useState(new Date());

  useEffect(() => {
    const timerId = setInterval(() => {
      setCurrentTime(new Date());
    }, 1000);
    return () => clearInterval(timerId);
  }, []);

  if (error) return <p className="error">Ошибка: {error}</p>;
  if (!weatherData) return null;

  return (
    <div className="weather-card">
      <h2>{weatherData.name}</h2>
      <div className="weather-info">
        <img 
          src={`http://openweathermap.org/img/wn/${weatherData.icon}@2x.png`} 
          alt="Weather icon"
        />
        <p>Температура: {weatherData.temperature}°C</p>
        <p>Время: {
          new Date(currentTime.getTime() + weatherData.timezone * 1000)
            .toISOString().substring(11, 19)
        }</p>
        <p>Описание: {weatherData.description}</p>
      </div>
    </div>
  );
}