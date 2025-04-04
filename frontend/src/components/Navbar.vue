<script setup>
import {RouterLink} from 'vue-router'
import {ref} from 'vue'

const cart = localStorage.getItem('cart')

const noOfItems = ref(JSON.parse(cart).length)

</script>

<template>
    <header>
        <h1>HOMEMADE-FOOD-MARKETPLACE</h1>
        <nav>
            <ul>
                <li><RouterLink to="/">Home</RouterLink></li>
                <li><RouterLink to="/menu">Menu</RouterLink></li>
                <li><RouterLink to="/login">Login</RouterLink></li>
                <li><RouterLink to="/cart">Cart <span id="cart-count" v-text="noOfItems"></span></RouterLink></li>
                <li><RouterLink to="/chefs">Chefs</RouterLink></li>
            </ul>
        </nav>   
    </header>
</template>