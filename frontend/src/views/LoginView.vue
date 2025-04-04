<template>
    <div class="form-container">
        <h2>Login to Your Account</h2>
        <form action="" method="post">
            <div class="input-group">
                <label for="email">Email</label>
                <input type="email" id="email" placeholder="Enter your email" required>
            </div>
            <div class="input-group">
                <label for="password">Password</label>
                <input type="password" id="password" placeholder="Enter your password" required>
            </div>
            <button type="submit" class="btn-submit">Login</button>
            <p>Don't have an account? <a href="signup.html">Sign Up</a></p>
        </form>
    </div>
</template>

<style scoped>
/* General Styles */
body {
    font-family: 'Poppins', sans-serif;
    background: linear-gradient(to right, #ffb347, #ffcc33);
    margin: 0;
    padding: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

/* Form Container */
.form-container {
    background: white;
    padding: 40px;
    border-radius: 10px;
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
    width: 380px;
    text-align: center;
    animation: fadeIn 0.5s ease-in-out;
}

.form-container h2 {
    margin-bottom: 20px;
    color: #d2691e;
    font-size: 26px;
}

/* Input Fields */
.input-group {
    margin-bottom: 20px;
    text-align: left;
}

.input-group label {
    display: block;
    font-weight: bold;
    margin-bottom: 5px;
    color: #333;
}

.input-group input {
    width: 100%;
    padding: 12px;
    border: 2px solid #d2691e;
    border-radius: 5px;
    font-size: 16px;
    outline: none;
    transition: all 0.3s ease-in-out;
}

.input-group input:focus {
    border-color: #ff9800;
    box-shadow: 0 0 8px rgba(255, 152, 0, 0.5);
}

/* Submit Button */
.btn-submit {
    width: 100%;
    background: #d2691e;
    color: white;
    padding: 14px;
    border: none;
    border-radius: 5px;
    font-size: 18px;
    font-weight: bold;
    cursor: pointer;
    transition: background 0.3s ease-in-out;
}

.btn-submit:hover {
    background: #b55d19;
}

/* Links */
.form-container p {
    margin-top: 15px;
    font-size: 14px;
    color: #555;
}

.form-container a {
    color: #d2691e;
    text-decoration: none;
    font-weight: bold;
}

.form-container a:hover {
    text-decoration: underline;
}

/* Responsive Design */
@media (max-width: 480px) {
    .form-container {
        width: 90%;
        padding: 30px;
    }
}

/* Fade-in Animation */
@keyframes fadeIn {
    from {
        opacity: 0;
        transform: translateY(-20px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}
</style>