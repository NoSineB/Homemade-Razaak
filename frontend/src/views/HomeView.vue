<script setup>
import cu from "@/assets/images/cu.jpg"
import images from "@/assets/images/images.jpg"
import women from "@/assets/images/women.jpg"
import chicken from "@/assets/images/chicken.jpg"
import pasta from "@/assets/images/pasta.jpg"
import dal from "@/assets/images/dal.jpg"

import Navbar from "@/components/Navbar.vue"
import Footer from "@/components/Footer.vue"

import {RouterLink} from 'vue-router'

</script>

<template>
  <div>


    <Navbar/>


    <section class="hero">
        <h2>Delicious Homemade Food, Just for You</h2>
        <p>Fresh, healthy, and prepared with love by talented housewives.</p>
        <RouterLink to="/menu" class="btn">Order Now</RouterLink>
    </section>


<section class="about">
  <p>We are dedicated to providing fresh, homemade meals prepared by skilled housewives in our community. Our mission is to support home chefs while delivering high-quality, nutritious food to busy individuals.</p>
    <div class="about-grid">
        <div class="about-item">
            <img :src="cu" alt="Home Cooked Food">
            <h3>100% Homemade</h3>
            <p>Every meal is prepared in a home kitchen with fresh ingredients.</p>
        </div>
        <div class="about-item">
            <img :src="images" alt="Healthy Food">
            <h3>Healthy & Tasty</h3>
            <p>We ensure a perfect balance of nutrition and taste in every bite.</p>
        </div>
        <div class="about-item">
            <img :src="women" alt="Support Women">
            <h3>Supporting Housewives</h3>
            <p>We empower women by helping them earn from their culinary skills.</p>
        </div>
    </div>
</section>


<section class="featured">
    <h2>Our Special Dishes</h2>
    <div class="menu">
        <div class="item">
            <img :src="chicken" alt="Biryani">
            <h3>Chicken Biryani</h3>
            <p>Spicy, flavorful, and cooked with authentic spices.</p>
        </div>
        <div class="item">
            <img :src="pasta" alt="Pasta">
            <h3>Cheesy Pasta</h3>
            <p>Rich and creamy homemade pasta with fresh herbs.</p>
        </div>
        <div class="item">
            <img :src="dal" alt="Dal Chawal">
            <h3>Dal Chawal</h3>
            <p>Traditional comfort food with a perfect blend of spices.</p>
        </div>
    </div>
    <RouterLink to="/menu" class="btn">Order Now</RouterLink>
</section>


<section class="testimonials">
    <h2>What Our Customers Say</h2>
    <div class="testimonial">
        <p>"The food is just like home-cooked meals. Healthy and delicious!"</p>
        <h4>- Ayesha K.</h4>
    </div>
    <div class="testimonial">
        <p>"Finally, a service that provides homemade food with amazing taste!"</p>
        <h4>- Rohit S.</h4>
    </div>
</section>


  <Footer />
  </div>
</template>

<style>
/* General Styling */
body {
    font-family: 'Poppins', sans-serif;
    background-color: #fffaf0;
    color: #333;
    margin: 0;
    padding: 0;
}

/* Header */
header {
    background: #d2691e;
    color: white;
    padding: 15px;
    text-align: center;
}

nav ul {
    list-style: none;
    padding: 0;
    display: flex;
    justify-content: center;
}

nav ul li {
    margin: 0 15px;
}

nav ul li a {
    color: white;
    text-decoration: none;
    font-weight: bold;
}

.hero {
    background: url('images/food-banner.jpg') no-repeat center center/cover;
    height: 300px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
    color: white;
}

.hero h2 {
    background: rgba(0, 0, 0, 0.6);
    padding: 10px;
}

.hero p {
    background: rgba(0, 0, 0, 0.4);
    padding: 5px;
    margin-top: 5px;
}

.btn {
    background: #ff9800;
    color: white;
    padding: 10px 20px;
    text-decoration: none;
    border-radius: 5px;
}

.about {
    text-align: center;
    padding: 40px;
    background: #f8f8f8;
}

.about-grid {
    display: flex;
    justify-content: center;
    gap: 20px;
}

.about-item {
    width: 30%;
    background: white;
    padding: 20px;
    border-radius: 10px;
}

.about-item img {
    width: 100%;
    border-radius: 10px;
}

.featured {
    text-align: center;
    padding: 40px;
}

.menu {
    display: flex;
    justify-content: center;
    gap: 20px;
}

.item {
    background: white;
    padding: 20px;
    border-radius: 10px;
}

.item img {
    width: 100%;
    border-radius: 10px;
}

.add-to-cart {
    background: #d2691e;
    color: white;
    padding: 10px;
    border: none;
    border-radius: 5px;
}

.testimonials {
    text-align: center;
    padding: 40px;
    background: #f8f8f8;
}

.testimonial {
    background: white;
    padding: 20px;
    border-radius: 10px;
    margin: 10px;
}

footer {
    background: #d2691e;
    color: white;
    text-align: center;
    padding: 10px;
}
/* General Styling */
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Poppins', sans-serif;
}

/* Body Styling */
body {
    background-color: #fffaf0;
    color: #333;
    line-height: 1.6;
}

/* Header */
header {
    background: #d2691e;
    color: white;
    padding: 15px 20px;
    text-align: center;
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

header h1 {
    font-size: 24px;
}

nav ul {
    list-style: none;
    display: flex;
    justify-content: center;
    margin-top: 10px;
}

nav ul li {
    margin: 0 15px;
}

nav ul li a {
    color: white;
    text-decoration: none;
    font-weight: bold;
    transition: color 0.3s ease;
}

nav ul li a:hover {
    color: #ffd700;
}

/* Hero Section */
.hero {
    background: url('images/food-banner.jpg') no-repeat center center/cover;
    height: 300px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
    color: white;
    font-size: 20px;
}

.hero h2 {
    background: rgba(0, 0, 0, 0.6);
    padding: 10px;
    border-radius: 5px;
}

.hero p {
    background: rgba(0, 0, 0, 0.4);
    padding: 5px;
    margin-top: 5px;
    border-radius: 5px;
}

.btn {
    display: inline-block;
    background: #ff9800;
    color: white;
    padding: 10px 20px;
    margin-top: 10px;
    text-decoration: none;
    border-radius: 5px;
    font-weight: bold;
    transition: background 0.3s;
}

.btn:hover {
    background: #e68900;
}

/* Menu Section */
.menu {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 20px;
    padding: 40px;
    text-align: center;
}

.item {
    background: white;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s ease;
}

.item:hover {
    transform: scale(1.05);
}

.item img {
    width: 100%;
    height: 200px;
    border-radius: 10px;
    object-fit: cover;
}

.item h3 {
    margin-top: 10px;
    font-size: 20px;
}

.item p {
    font-size: 14px;
    color: #777;
}

.add-to-cart {
    background: #d2691e;
    color: white;
    padding: 10px;
    border: none;
    cursor: pointer;
    margin-top: 10px;
    font-weight: bold;
    border-radius: 5px;
    transition: background 0.3s;
}

.add-to-cart:hover {
    background: #a0522d;
}

/* About Housewives Section */
.chefs {
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    padding: 40px;
    gap: 20px;
}

.chef {
    background: white;
    padding: 20px;
    border-radius: 10px;
    text-align: center;
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
    width: 300px;
}

.chef img {
    width: 100%;
    height: 200px;
    border-radius: 10px;
    object-fit: cover;
}

.chef h3 {
    margin-top: 10px;
    font-size: 18px;
}

.chef p {
    font-size: 14px;
    color: #777;
}

/* Checkout Page */
.checkout {
    text-align: center;
    padding: 40px;
}

#google-pay {
    background: #34a853;
    color: white;
    padding: 15px 20px;
    border: none;
    cursor: pointer;
    font-size: 16px;
    font-weight: bold;
    border-radius: 5px;
    transition: background 0.3s;
}

#google-pay:hover {
    background: #2a8c45;
}

/* Footer */
footer {
    background: #d2691e;
    color: white;
    text-align: center;
    padding: 10px;
    margin-top: 20px;
} 
</style>
