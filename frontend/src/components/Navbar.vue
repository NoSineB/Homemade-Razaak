<script setup>
import axios from 'axios';
import { onMounted, ref } from 'vue';

const userId = sessionStorage.getItem('user_id') ? sessionStorage.getItem('user_id') : 0
const noOfItems = ref()

onMounted(async function() { 
    try{
        const response = await axios.post('http://localhost:3000/cart/len', {
           user_id: parseInt(userId)
        })
        noOfItems.value = response.data.length

    }catch (error) {
        noOfItems.value = ref()
    }
    }
)

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