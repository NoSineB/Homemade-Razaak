<script setup>
    import Navbar from "@/components/Navbar.vue"
    import Footer from "@/components/Footer.vue"

    import meal from "@/assets/images/meal.jpg"
    import chicken from "@/assets/images/chicken.jpg"
    import fish from "@/assets/images/fish.jpg"
    import kanji from "@/assets/images/kanji.jpg"
    import fishcurry from "@/assets/images/fishcurry.jpg"
    import achar from "@/assets/images/achar.jpg"
    import Chappathi from "@/assets/images/Chappathi.jpg"
    import beafCurry from "@/assets/images/beaf curry.jpg"

    import axios from "axios"

    const menuItems = [
        {
            id:1,
            itemName: "Meals",
            image: meal,
            description: "Delicious Homemade Meals",
            price: 150
        },
        {
            id:2,
            itemName: "Chicken Briyani",
            image: chicken,
            description: "Authentic biryani with homemade spices.",
            price: 250
        },
        {
            id:3,
            itemName: "Fish Fry",
            image: fish,
            description: "Crispy and spicy homemade fishfry .",
            price: 75
        },
        {
            id:4,
            itemName: "Kanji",
            image: kanji,
            description: "Flowy delicious homemade kanji.",
            price: 50
        },
        {
            id:5,
            itemName: "Fish Curry",
            image: fishcurry,
            description: "Softy fish with gravy.",
            price: 350
        },
        {
            id:6,
            itemName: "Achar",
            image: achar,
            description: "Delicious homemade achar.",
            price: 110
        },
        {
            id:7,
            itemName: "Chappathi",
            image: Chappathi,
            description: "Soft and fluffy home made chappathi.",
            price: 20
        },
        {
            id:8,
            itemName: "Beaf Curry",
            image: beafCurry,
            description: "Rich and flavoured beef homemade curry.",
            price: 110
        },
    ]

async function addToCart(id) {
    const user_id = sessionStorage.getItem('user_id') ? sessionStorage.getItem('user_id') : 0

    try{
        const response = await axios.post('http://localhost:3000/cart', {
            user_id: parseInt(user_id),
            item_id: parseInt(id)
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

    <section class="menu">
        <div class="item" v-for="each in menuItems" :key="each.id">
            <img :src="each.image" :alt="each.itemName">
            <h3 v-text="each.itemName"></h3>
            <p v-text="each.description"></p>
            <p><strong>₹ <span v-text="each.price"></span></strong></p>
            <button class="btn add-to-cart" @click="addToCart(each.id)">Add to Cart</button>
        </div>
    </section>

    <Footer />
</template>