<script>
    import Image from '$lib/assets/download.png'
    let name="Meles"
    let count=0;
    const increment =()=>{
        count+=1;
    }

    $:double=count*2
    $:console.log(`the updated value is ${double}`)
    if (count>5){
        alert("Count is not greater than five")
    }
    let numbers=[1,3,4]

    function addNewNumbers(){
        numbers=[...numbers, numbers.length+1]
        console.log(numbers)
    }


</script>
<p>The numbers are {numbers}</p>

<button on:click={addNewNumbers}>Add numbers</button>
<p>Increase the value to {count}</p>

<p>Double my count is {double}</p>

<button on:click={increment}>Increase</button>

<p>My name is {name}</p>
<img src={Image} alt="Here we are">