import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import WeatherForm from '../components/WeatherForm';
import WeatherDisplay from '../components/WeatherDisplay';

export default function CityPage() {
  const { cityName } = useParams();
  const [weatherData, setWeatherData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(`/api/getWeather/${encodeURIComponent(cityName)}`);
        if (!response.ok) throw new Error('Город не найден');
        const data = await response.json();
        setWeatherData(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [cityName]);

  return (
    <div className="app">
      <h1>weather.time</h1>
      <WeatherForm />
      {loading ? <p>Загрузка...</p> : <WeatherDisplay weatherData={weatherData} error={error} />}
    </div>
  );
}