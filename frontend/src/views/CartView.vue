<script setup>
import Footer from "@/components/Footer.vue"
import { onMounted, ref } from "vue"
import { useRouter } from "vue-router"
import axios from "axios"
import Navbar from "@/components/Navbar.vue"

const total = ref()
const router = useRouter()

onMounted(async function() { 
    const user_id = sessionStorage.getItem('user_id') ? sessionStorage.getItem('user_id') : 0

    try{
        const response = await axios.post('http://localhost:3000/cart/total', {
        user_id: parseInt(user_id)
        })
        total.value = response.data.total

    }catch (error) {
        console.error(error)
        router.push({path: "/login"})
    }
})

async function checkout() {
    const user_id = sessionStorage.getItem('user_id') ? sessionStorage.getItem('user_id') : 0

    try{
        const response = await axios.post('http://localhost:3000/cart/checkout', {
            user_id: parseInt(user_id),
        })

        console.log(response.data)

    }catch (error) {
        console.error(error)
        router.push({path: "/login"})
    }
}

</script>

<template>
    <Navbar/>
    <header>
        <h1>Checkout</h1>
    </header>

    <section class="checkout">
        <h2>Your Cart</h2>
        <ul id="cart-items"></ul>
        <p>Total: ₹<span id="cart-total" v-text="total"></span></p>
        
        <button id="google-pay" @click="checkout">Pay with Google Pay</button>
    </section>
    <Footer />
</template>