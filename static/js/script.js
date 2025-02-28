// script.js
document.addEventListener('DOMContentLoaded', function() {
    // Обработчик для кнопки "Полное описание"
    const fullDescBtn = document.querySelector('.btn-full-description');
    const fullDescSection = document.getElementById('full-description-section');
    if (fullDescBtn && fullDescSection) {
        fullDescBtn.addEventListener('click', function() {
            fullDescSection.scrollIntoView({ behavior: 'smooth' });
        });
    }

    const allCharactBtn = document.querySelector('.btn-all-characteristics');
    const characteristicsSection = document.getElementById('characteristics-section');
    if (allCharactBtn && characteristicsSection) {
        allCharactBtn.addEventListener('click', function() {
            characteristicsSection.scrollIntoView({ behavior: 'smooth' });
        });
    }

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
        }, { threshold: 0.1 });

        animateElements.forEach(el => {
            observer.observe(el);
        });
    }
});