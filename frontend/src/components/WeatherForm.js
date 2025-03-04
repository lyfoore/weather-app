import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function WeatherForm() {
  const [city, setCity] = useState('');
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    if (city.trim()) {
      navigate(`/${encodeURIComponent(city.trim())}`);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={city}
        onChange={(e) => setCity(e.target.value)}
        placeholder="Введите город"
      />
      <button type="submit">Поиск</button>
    </form>
  );
}