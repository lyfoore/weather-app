import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import WeatherForm from '../components/WeatherForm';
import WeatherDisplay from '../components/WeatherDisplay';

export default function HomePage() {
  const navigate = useNavigate();

  // Автоматическая загрузка Новосибирска при монтировании
  useEffect(() => {
    navigate('/Novosibirsk');
  }, [navigate]);

  return (
    <div className="app">
      <h1>weather.time</h1>
      <WeatherForm />
      <WeatherDisplay />
    </div>
  );
}