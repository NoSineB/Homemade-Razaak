<script setup>
    import Footer from "@/components/Footer.vue"

    const cart = localStorage.getItem('cart')SON.parse(cart))
</script>

<template>
    <header>
        <h1>Checkout</h1>
    </header>

    <section class="checkout">
        <h2>Your Cart</h2>
        <ul id="cart-items"></ul>
        <p>Total: ₹<span id="cart-total">0</span></p>
        
        <button id="google-pay">Pay with Google Pay</button>
    </section>