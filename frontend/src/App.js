import React, { useState } from 'react';
import './App.css';

function App() {
  const [city, setCity] = useState('');
  const [weatherData, setWeatherData] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchWeather = async (e) => {
    e.preventDefault();
    if (!city) return;
    
    setLoading(true);
    setError(null);
    
    try {
      const response = await fetch(`/api/getResponse/${encodeURIComponent(city)}`);
      if (!response.ok) {
        throw new Error('Город не найден');
      }
      const data = await response.json();
      setWeatherData(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="app">
      <h1>weather.time</h1>
      
      <form onSubmit={fetchWeather}>
        <input
          type="text"
          value={city}
          onChange={(e) => setCity(e.target.value)}
          placeholder="Введите город"
        />
        <button type="submit">Поиск</button>
      </form>

      {loading && <p>Загрузка...</p>}
      
      {error && <p className="error">Ошибка: {error}</p>}

      {weatherData && (
        <div className="weather-card">
          <h2>{weatherData.name}</h2>
          <div className="weather-info">
            <img 
              src={`http://openweathermap.org/img/wn/${weatherData.icon}@2x.png`} 
              alt="Weather icon"
            />
            <p>Температура: {weatherData.temperature}°C</p>
            <p>Время: {weatherData.time}</p>
            <p>Описание: {weatherData.description}</p>
          </div>
        </div>
      )}
    </div>
  );
}

export default App;