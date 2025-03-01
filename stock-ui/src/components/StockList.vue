<template>
  <div class="container mx-auto p-6 bg-gray-50 min-h-screen">
    <h1 class="text-3xl font-bold mb-8 text-center text-gray-800">Stock List</h1>

    <!--Botones para diferentes fetchs -->
    <div class="flex justify-center mb-8">
      <button 
        @click="fetchStocks" 
        class="bg-blue-500 hover:bg-blue-600 text-white font-semibold px-4 py-2 rounded-lg mr-4"
      >
        Fetch Stocks
      </button>
      <button 
        @click="fetchBestStocks"
        class="bg-green-500 hover:bg-green-600 text-white font-semibold px-4 py-2 rounded-lg"
      >
        Fetch Best Stocks
      </button>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="stock in stocks" 
        :key="stock.ticker" 
        class="bg-white rounded-xl shadow-md overflow-hidden hover:shadow-lg transition duration-300 transform hover:-translate-y-1"
      >

        <div 
          :class="[
            'p-4 text-white', 
            getStockColor(stock.rating_from, stock.rating_to)
          ]"
        >
          <div class="flex justify-between items-center">
            <h2 class="text-xl font-bold truncate">{{ stock.company }}</h2>
            <span class="font-mono text-white bg-white bg-opacity-20 px-2 py-1 rounded text-sm">
              {{ stock.ticker }}
            </span>
          </div>
        </div>
        
        <div class="p-4">
          <div class="grid grid-cols-2 gap-4 mb-4">
            <div>
              <div class="text-xs text-gray-500 uppercase font-semibold">Brokerage</div>
              <div class="font-medium">{{ stock.brokerage }}</div>
            </div>
            <div>
              <div class="text-xs text-gray-500 uppercase font-semibold">Action</div>
              <div 
                :class="[
                  'font-medium rounded-full px-2 py-1 text-center text-sm', 
                  getStockColor(stock.rating_from, stock.rating_to) === 'bg-green-600' ? 'bg-green-100 text-green-600' :
                  getStockColor(stock.rating_from, stock.rating_to) === 'bg-blue-600' ? 'bg-blue-100 text-blue-600' : 'bg-red-100 text-red-600'
                ]"
              >
                {{ stock.action }}
              </div>
            </div>
          </div>

          <div class="space-y-2 mb-4">
            <div>
              <div class="text-xs text-gray-500 uppercase font-semibold">Rating</div>
              <div class="flex items-center justify-center">
                <span class="text-gray-600">{{ stock.rating_from }}</span>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mx-2 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                </svg>
                <span class="font-medium">{{ stock.rating_to }}</span>
              </div>
            </div>
            
            <div>
              <div class="text-xs text-gray-500 uppercase font-semibold">Target</div>
              <div class="flex items-center justify-center">
                <span class="text-gray-600">${{ stock.target_from }}</span>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mx-2 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                </svg>
                <span class="font-medium">${{ stock.target_to }}</span>
              </div>
            </div>
          </div>

          <div class="text-xs text-gray-500 mt-4 flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            {{ formatDate(stock.time) }}
          </div>
        </div>
      </div>
    </div>

    <div v-if="stocks.length === 0" class="text-center py-12 text-gray-500">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto mb-4 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
      </svg>
      <p class="text-lg">No stocks available at the moment</p>
    </div>
  </div>
</template>

<script>
const BASE_API_URL = 'http://localhost:8080/api'

export default {
  data() {
    return {
      stocks: []
    }
  },
  async created() {
    this.fetchStocks()
  },
  methods: {
    async fetchStocks() {
      try {
        const response = await fetch(`${BASE_API_URL}/stocks`)
        const data = await response.json()
        this.stocks = data
      } catch (error) {
        console.error('Error fetching stocks:', error)
      }
    },
    async fetchBestStocks() {
      try {
        const response = await fetch(`${BASE_API_URL}/stocks/recommendations`)
        const data = await response.json()
        this.stocks = data
      } catch (error) {
        console.error('Error fetching best stocks:', error)
      }
    },
    getStockColor(from, to) {
      const positiveChanges = {
        'Sell': ['Neutral', 'Buy', 'Outperform', 'Strong Buy'],
        'Neutral': ['Buy', 'Outperform'],
        'Underweight': ['Equal Weight', 'Overweight', 'Buy'],
        'Equal Weight': ['Overweight', 'Buy'],
        'Market Perform': ['Outperform', 'Buy'],
        'Hold': ['Buy', 'Strong Buy'],
        'Buy': ['Overweight', 'Strong Buy', 'Outperform', 'Buy'],
      }

      if (positiveChanges[from] && positiveChanges[from].includes(to)) {
        return 'bg-green-500'
      }

      if (from == to) {
        return 'bg-gray-500'
      }

      return 'bg-red-500'
    },
    formatDate(date) {
      return new Date(date).toLocaleString()
    }
  }
}
</script>
