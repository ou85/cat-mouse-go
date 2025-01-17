const leaderboardData = [
    { name: 'Twisted Mind', score: 150 },
    { name: 'Teddy the Bear', score: 120 },
    { name: 'Shine vombat', score: 80 },
    { name: 'Speedy Gonzales', score: 110 },
    { name: 'Fromage Français', score: 95 },
    { name: 'Sauerkraut Fritz', score: 105 },
    { name: 'Pizza Pete', score: 130 },
    { name: 'Coco Chanel', score: 85 },
    { name: 'Rocky Balboa', score: 125 },
    { name: 'Sunny Day', score: 90 }
  ];
  
  // Функция для рендеринга таблицы
  function renderLeaderboard(data) {
    const tbody = document.getElementById('leaderboard-body');
    tbody.innerHTML = ''; // Очистка старых данных
  
    data.forEach((entry, index) => {
      const row = document.createElement('tr');
      row.innerHTML = `
        <td>${index + 1}</td>
        <td>${entry.name}</td>
        <td>${entry.score}</td>
      `;
      tbody.appendChild(row);
    });
  }
  
  // Функция для обновления результатов
  function updateScore(name, score) {
    leaderboardData.push({ name, score });
    leaderboardData.sort((a, b) => b.score - a.score); // Сортировка по убыванию
    leaderboardData.splice(10); // Хранить только топ-10
    renderLeaderboard(leaderboardData); // Обновление таблицы
  }
  
  // Инициализация
  leaderboardData.sort((a, b) => b.score - a.score); // Первичная сортировка
  renderLeaderboard(leaderboardData); // Рендеринг таблицы
  
  // Пример добавления нового результата
  updateScore('Fast Fox', 140);
  