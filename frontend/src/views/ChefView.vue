<script setup>
import meen from "@/assets/images/meen.jpg"
import teen from "@/assets/images/teen.jpg"
import gf from "@/assets/images/gf.jpg"

import Footer from "@/components/Footer.vue"

import {RouterLink} from 'vue-router'
</script>

<template>
    <header>
        <h1>Meet the Chefs</h1>
        <nav>
            <ul>
                <li><RouterLink to="/">Home</RouterLink></li>
                <li><RouterLink to="/menu">Menu</RouterLink></li>
                <li><RouterLink to="/cart">Cart</RouterLink></li>
            </ul>
        </nav>
    </header>

    <section class="chefs">
        <div class="chef">
            <img :src="meen" alt="Housewife 1">
            <h3>Mrs. Aisha</h3>
            <p>Specialist in Biryani and Traditional Curries.</p>
        </div>

        <div class="chef">
            <img :src="teen" alt="Housewife 2">
            <h3>Mrs. Meena</h3>
            <p>Master of Homemade Parathas and Sweets.</p>
        </div>

        <div class="chef">
            <img :src="gf" alt="Housewife 3">
            <h3>priya</h3>
            <p>Specialist in achar and Traditional Curries.</p>
        </div>
    </section>

    <Footer/>
</template>