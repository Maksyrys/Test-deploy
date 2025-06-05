document.addEventListener('DOMContentLoaded', function() {
    /**
     * Универсальная функция для вывода ошибки в контейнер с указанным id
     * Если контейнера на странице нет, просто выводим ошибку в консоль (fallback).
     */
    function showError(containerId, message) {
        const el = document.getElementById(containerId);
        if (el) {
            el.textContent = message;
            el.style.display = 'block';  // Показываем блок
        } else {
            console.error("Ошибка (контейнер не найден):", message);
        }
    }

    /** Скрыть ошибку в конкретном контейнере (например, при повторных запросах) */
    function hideError(containerId) {
        const el = document.getElementById(containerId);
        if (el) {
            el.style.display = 'none';
            el.textContent = '';
        }
    }

    /** "Глобальный" вывод ошибки, если нет особого контейнера для конкретной операции */
    function showGlobalError(message) {
        showError('globalError', message);
    }

    // --- Функция обновления счётчика товаров в корзине ---
    function updateCartIconCount() {
        fetch('/cart/count', { method: 'GET' })
            .then(response => response.json())
            .then(data => {
                const cartCountEl = document.getElementById('cart-count');
                if (cartCountEl) {
                    cartCountEl.innerText = data.count;
                }
            })
            .catch(err => console.error('Ошибка обновления корзины:', err));
    }

    // --- Функция пересчёта итоговой суммы и количества позиций в корзине ---
    function updateCartSummary() {
        let grandTotal = 0;
        let totalQuantity = 0;

        document.querySelectorAll('.cart-item-row').forEach(item => {
            const quantityEl = item.querySelector('.item-quantity');
            const quantity = parseInt(quantityEl.innerText);

            // Из текста цены убираем всё лишнее
            const priceText = item.querySelector('.item-price').innerText
                .replace('Цена:', '')
                .replace('₽', '')
                .trim();
            const price = parseFloat(priceText);

            const rowTotal = quantity * price;

            // Если есть ячейка с суммой по этой позиции, обновим
            const rowTotalEl = item.querySelector('.cart-item-total');
            if (rowTotalEl) {
                rowTotalEl.innerText = rowTotal.toFixed(2) + ' ₽';
            }

            grandTotal += rowTotal;
            totalQuantity += quantity;
        });

        const grandTotalEl = document.getElementById('grand-total');
        const totalQuantityEl = document.getElementById('total-quantity');
        if (grandTotalEl) grandTotalEl.innerText = grandTotal.toFixed(2) + ' ₽';
        if (totalQuantityEl) totalQuantityEl.innerText = totalQuantity;
    }

    // ----------------------
    //  Прокрутки к секциям
    // ----------------------
    updateCartIconCount();

    const fullDescBtn = document.querySelector('.btn-full-description');
    const fullDescSection = document.getElementById('full-description-section');
    if (fullDescBtn && fullDescSection) {
        fullDescBtn.addEventListener('click', function () {
            fullDescSection.scrollIntoView({behavior: 'smooth'});
        });
    }

    const allCharactBtn = document.querySelector('.btn-all-characteristics');
    const characteristicsSection = document.getElementById('characteristics-section');
    if (allCharactBtn && characteristicsSection) {
        allCharactBtn.addEventListener('click', function () {
            characteristicsSection.scrollIntoView({behavior: 'smooth'});
        });
    }

    // ----------------------
    // Анимация при скролле
    // ----------------------
    const animateElements = document.querySelectorAll('.animate-on-scroll');
    if (animateElements.length > 0) {
        const observer = new IntersectionObserver((entries) => {
            entries.forEach(entry => {
                if (entry.isIntersecting) {
                    entry.target.classList.add('fade-in');
                } else {
                    entry.target.classList.remove('fade-in');
                }
            });
        }, {threshold: 0.1});
        animateElements.forEach(el => observer.observe(el));
    }

    // ----------------------
    //  Добавление в корзину
    // ----------------------
    const addToCartBtn = document.querySelector('.btn-add-cart');
    if (addToCartBtn) {
        addToCartBtn.addEventListener('click', async function () {

            if (this.dataset.inCart === 'true') {
                window.location.href = '/cart';
                return;
            }

            const bookId = this.getAttribute('data-book-id');
            const formData = new URLSearchParams();
            formData.append('book_id', bookId);

            try {
                const response = await fetch('/cart/add', {
                    method: 'POST',
                    body: formData,
                    headers: {"Content-Type": "application/x-www-form-urlencoded"}
                });

                if (!response.ok) {
                    if (response.status === 401) {
                        // Если не авторизован – предлагаем логин
                        showLoginPrompt();
                        return;
                    } else {
                        const errorText = await response.text();
                        showGlobalError("Ошибка при добавлении в корзину: " + errorText);
                        return;
                    }
                }

                const result = await response.json();
                if (result.success) {
                    this.dataset.inCart = 'true';
                    this.innerText = 'Перейти в корзину';
                    updateCartIconCount();
                } else {
                    showGlobalError("Ошибка добавления в корзину");
                }
            } catch (err) {
                console.error('Ошибка:', err);
            }
        });
    }

    // Если пользователь кликает по кнопке, но не авторизован (cart, favorites)
    const unauthCart = document.getElementById('unauthCart');
    if (unauthCart) {
        unauthCart.addEventListener('click', function (e) {
            e.preventDefault();
            showLoginPrompt();
        });
    }

    const unauthFavorites = document.getElementById('unauthFavorites');
    if (unauthFavorites) {
        unauthFavorites.addEventListener('click', function(e) {
            e.preventDefault();
            showLoginPrompt();
        });
    }

    // ----------------------
    // Модальное окно аутентификации
    // ----------------------
    const openAuthModalBtn = document.getElementById("openAuthModalBtn");
    const authModal = document.getElementById("authModal");
    const closeAuthModalBtn = document.getElementById("closeAuthModalBtn");
    const authTitle = document.getElementById("authTitle");
    const authButtons = document.getElementById("authButtons");
    const showLoginFormBtn = document.getElementById("showLoginFormBtn");
    const showRegisterFormBtn = document.getElementById("showRegisterFormBtn");
    const loginForm = document.getElementById("loginForm");
    const registerForm = document.getElementById("registerForm");

    if (openAuthModalBtn && authModal && closeAuthModalBtn) {
        openAuthModalBtn.addEventListener("click", function (e) {
            e.preventDefault();
            authModal.style.display = "block";
            authTitle.style.display = "block";
            authButtons.style.display = "flex";
            if (loginForm) loginForm.style.display = "none";
            if (registerForm) registerForm.style.display = "none";
        });

        closeAuthModalBtn.addEventListener("click", function () {
            authModal.style.display = "none";
        });

        if (showLoginFormBtn) {
            showLoginFormBtn.addEventListener("click", function () {
                authTitle.style.display = "none";
                authButtons.style.display = "none";
                if (loginForm) loginForm.style.display = "block";
                if (registerForm) registerForm.style.display = "none";
            });
        }

        if (showRegisterFormBtn) {
            showRegisterFormBtn.addEventListener("click", function () {
                authTitle.style.display = "none";
                authButtons.style.display = "none";
                if (loginForm) loginForm.style.display = "none";
                if (registerForm) registerForm.style.display = "block";
            });
        }

        window.addEventListener("click", function (event) {
            if (event.target === authModal) {
                authModal.style.display = "none";
            }
        });
    }

    // ----------------------
    //   Вход (логин)
    // ----------------------
    if (loginForm) {
        loginForm.addEventListener('submit', async function (e) {
            e.preventDefault();

            // Скрываем предыдущую ошибку (если была)
            hideError('loginError');

            const formDataObj = new URLSearchParams(new FormData(this));

            try {
                const response = await fetch('/login', {
                    method: 'POST',
                    body: formDataObj,
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    },
                    credentials: 'same-origin'
                });

                if (!response.ok) {
                    const errorText = await response.text();
                    showError('loginError', errorText);  // Выводим ошибку в контейнер формы логина
                    return;
                }

                const result = await response.json();
                if (result.success) {
                    location.reload();
                } else {
                    showError('loginError', "Ошибка входа");
                }
            } catch (err) {
                console.error('Ошибка:', err);
                showError('loginError', "Не удалось отправить запрос на вход.");
            }
        });
    }

    // ----------------------
    //   Регистрация
    // ----------------------
    if (registerForm) {
        registerForm.addEventListener('submit', async function (e) {
            e.preventDefault();

            hideError('registerError');

            // Кастомная валидация для пароля
            const password = document.getElementById('reg-password').value;
            if (password.length < 3) {
                showError('registerError', 'Пароль должен содержать минимум 3 символа');
                return;
            }

            const formDataObj = new URLSearchParams(new FormData(this));
            try {
                const response = await fetch('/register', {
                    method: 'POST',
                    body: formDataObj,
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    },
                    credentials: 'same-origin'  // Вот тут!
                });

                if (!response.ok) {
                    const errorText = await response.text();
                    showError('registerError', errorText);
                    return;
                }

                const result = await response.json();
                if (result.success) {
                    location.reload();
                } else {
                    showError('registerError', "Ошибка регистрации");
                }
            } catch (err) {
                console.error('Ошибка:', err);
                showError('registerError', "Не удалось отправить запрос на регистрацию.");
            }
        });
    }


    // ----------------------
    //  Выход из системы
    // ----------------------
    const logoutBtn = document.getElementById('logoutBtn');
    if (logoutBtn) {
        logoutBtn.addEventListener('click', function (e) {
            e.preventDefault();
            window.location.href = '/logout'; // Просто перенаправляем — сервер всё сам сделает
        });
    }

    // ----------------------
    //   Корзина
    // ----------------------
    // Удаление товара из корзины
    document.querySelectorAll('.btn-remove-item').forEach(btn => {
        btn.addEventListener('click', async function () {
            const bookId = this.getAttribute('data-book-id');
            const formData = new URLSearchParams();
            formData.append('book_id', bookId);

            try {
                const response = await fetch('/cart/remove/all', {
                    method: 'POST',
                    body: formData,
                    headers: {"Content-Type": "application/x-www-form-urlencoded"}
                });
                if (!response.ok) {
                    const errorText = await response.text();
                    showGlobalError("Ошибка при удалении: " + errorText);
                    return;
                }
                const result = await response.json();
                if (result.success) {
                    location.reload();
                } else {
                    showGlobalError("Ошибка при удалении товара");
                }
            } catch (err) {
                console.error('Ошибка:', err);
            }
        });
    });

    // Оформление заказа
    const checkoutBtn = document.querySelector('.btn-checkout');
    if (checkoutBtn) {
        checkoutBtn.addEventListener('click', async function () {
            try {
                const response = await fetch('/cart/checkout', {method: 'POST'});
                if (!response.ok) {
                    const errorText = await response.text();
                    showGlobalError("Ошибка при оформлении заказа: " + errorText);
                    return;
                }
                window.location.href = '/order/confirmation';
            } catch (err) {
                console.error('Ошибка:', err);
            }
        });
    }

    // Увеличение количества товара
    document.querySelectorAll('.btn-increase').forEach(btn => {
        btn.addEventListener('click', async function () {
            const bookId = this.getAttribute('data-book-id');
            const formData = new URLSearchParams();
            formData.append('book_id', bookId);
            formData.append('quantity', '1');

            try {
                const response = await fetch('/cart/add', {
                    method: 'POST',
                    body: formData,
                    headers: {"Content-Type": "application/x-www-form-urlencoded"}
                });
                if (!response.ok) {
                    const errorText = await response.text();
                    showGlobalError("Ошибка при увеличении количества: " + errorText);
                    return;
                }
                const quantityEl = this.parentElement.querySelector('.item-quantity');
                const newQuantity = parseInt(quantityEl.innerText) + 1;
                quantityEl.innerText = newQuantity;
                updateCartSummary();
                updateCartIconCount();
            } catch (err) {
                console.error('Ошибка:', err);
            }
        });
    });

    // Уменьшение количества товара
    document.querySelectorAll('.btn-decrease').forEach(btn => {
        btn.addEventListener('click', async function () {
            const bookId = this.getAttribute('data-book-id');
            const formData = new URLSearchParams();
            formData.append('book_id', bookId);

            try {
                const response = await fetch('/cart/remove', {
                    method: 'POST',
                    body: formData,
                    headers: {"Content-Type": "application/x-www-form-urlencoded"}
                });
                if (!response.ok) {
                    const errorText = await response.text();
                    showGlobalError("Ошибка при уменьшении количества: " + errorText);
                    return;
                }
                const quantityEl = this.parentElement.querySelector('.item-quantity');
                let currentQuantity = parseInt(quantityEl.innerText);
                if (currentQuantity > 1) {
                    quantityEl.innerText = currentQuantity - 1;
                } else {
                    // Удаляем целиком блок товара
                    this.closest('.cart__item').remove();
                }
                updateCartSummary();
                updateCartIconCount();
            } catch (err) {
                console.error('Ошибка:', err);
            }
        });
    });

    // ----------------------
    //   Избранное
    // ----------------------
    const favoriteButtons = document.querySelectorAll('.btn-favorite');
    favoriteButtons.forEach(btn => {
        btn.addEventListener('click', async function () {
            const bookId = btn.dataset.bookId;
            const isActive = btn.classList.contains('active');
            const url = isActive ? '/favorites/remove' : '/favorites/add';

            try {
                const response = await fetch(url, {
                    method: 'POST',
                    headers: {'Content-Type': 'application/x-www-form-urlencoded'},
                    body: `book_id=${bookId}`
                });
                if (!response.ok) {
                    if (response.status === 401) {
                        showLoginPrompt();
                        return;
                    }
                    const errorText = await response.text();
                    showGlobalError("Ошибка избранного: " + errorText);
                    return;
                }
                const data = await response.json();
                if (data.success) {
                    btn.classList.toggle('active');
                } else {
                    showGlobalError("Не удалось изменить избранное");
                }
            } catch (error) {
                console.error('Ошибка:', error);
            }
        });
    });

    // ----------------------
    //   Модальное окно «Войти или Зарегистрироваться»
    //   (showLoginPrompt) – при 401 ошибках
    // ----------------------
    function showLoginPrompt() {
        const modalContainer = document.createElement('div');
        modalContainer.classList.add('modal');
        modalContainer.id = 'loginPromptModal';
        modalContainer.style.display = 'block';
        modalContainer.innerHTML = `
            <div class="modal-content">
                <span class="close" id="closeLoginPrompt">&times;</span>
                <p>Чтобы использовать корзину или избранное, вам необходимо войти или зарегистрироваться на сайте.<br>Хотите это сделать?</p>
                <div style="text-align: center; margin-top: 20px;">
                    <button id="loginPromptYes" class="btn">Да</button>
                    <button id="loginPromptNo" class="btn">Нет</button>
                </div>
            </div>
        `;
        document.body.appendChild(modalContainer);

        document.getElementById('loginPromptYes').addEventListener('click', function() {
            const authModal = document.getElementById('authModal');
            if (authModal) {
                authModal.style.display = 'block';
            }
            closeLoginPrompt();
        });
        document.getElementById('loginPromptNo').addEventListener('click', closeLoginPrompt);
        document.getElementById('closeLoginPrompt').addEventListener('click', closeLoginPrompt);

        function closeLoginPrompt() {
            modalContainer.style.display = 'none';
            document.body.removeChild(modalContainer);
        }
    }

    // ----------------------
    //   Отзывы (Review)
    // ----------------------
    const reviewForm = document.getElementById('review-form');
    if (reviewForm) {
        const ratingStars = document.querySelectorAll('.stars-rating span');
        const commentEl = document.getElementById('review-comment');
        const bookIdEl = document.getElementById('review-book-id');
        let currentRating = 0;

        if (bookIdEl) {
            const bookId = bookIdEl.value;

            // Вешаем выбор звёзд
            ratingStars.forEach(star => {
                star.addEventListener('click', function () {
                    ratingStars.forEach(s => {
                        s.classList.toggle('active', s.dataset.value <= this.dataset.value);
                    });
                    currentRating = parseInt(this.dataset.value);
                });
            });

            // Отправка формы отзыва
            reviewForm.addEventListener('submit', async function (e) {
                e.preventDefault();
                hideError('reviewError'); // Предполагаем, что есть контейнер с id="reviewError"

                const comment = (commentEl.value || '').trim();
                if (!currentRating) {
                    showError('reviewError', 'Пожалуйста, поставьте оценку.');
                    return;
                }

                const formData = new URLSearchParams();
                formData.append('book_id', bookId);
                formData.append('rating', currentRating);
                formData.append('comment', comment);

                try {
                    const response = await fetch('/reviews/add', {
                        method: 'POST',
                        headers: {"Content-Type": "application/x-www-form-urlencoded"},
                        body: formData
                    });
                    if (!response.ok) {
                        const errorText = await response.text();
                        showError('reviewError', "Ошибка отправки отзыва: " + errorText);
                        return;
                    }
                    const result = await response.json();
                    if (result.success) {
                        location.reload();
                    } else {
                        showError('reviewError', "Ошибка отправки отзыва");
                    }
                } catch (err) {
                    console.error('Ошибка:', err);
                    showError('reviewError', "Не удалось отправить отзыв.");
                }
            });
        }
    }

    // ----------------------
    //   Отображение звёзд рейтинга (если есть)
    // ----------------------
    const starsRating = document.querySelectorAll('.stars-rating span');
    let currentRating = 0;
    starsRating.forEach(star => {
        star.addEventListener('click', function () {
            currentRating = parseInt(this.dataset.value);
            starsRating.forEach(s => {
                s.classList.toggle('active', parseInt(s.dataset.value) <= currentRating);
            });
        });
    });

    // Отрисовка статических звёзд, напр. на карточке, где уже есть рейтинг
    document.querySelectorAll('.stars-display').forEach(el => {
        const rating = parseInt(el.dataset.rating);
        // Простейший вариант: заполняем часть звёзд "★", а остальные "☆"
        el.innerHTML = '★★★★★'.slice(0, rating) + '☆☆☆☆☆'.slice(rating);
    });
});

document.addEventListener('DOMContentLoaded', function() {
    document.querySelectorAll('.admin-sidebar ul li').forEach(item => {
        item.addEventListener('click', () => {
            // Сброс активного класса у всех элементов боковой панели
            document.querySelectorAll('.admin-sidebar ul li').forEach(li => li.classList.remove('active'));
            item.classList.add('active');

            // Скрыть все секции
            document.querySelectorAll('.admin-content .admin-section').forEach(section => {
                section.style.display = 'none';
            });

            // Показать выбранную секцию
            const target = item.getAttribute('data-target');
            document.getElementById(target).style.display = 'block';
        });
    });
});