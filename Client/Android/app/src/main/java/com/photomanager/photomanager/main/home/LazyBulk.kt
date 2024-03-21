package com.photomanager.photomanager.main.home

import android.util.LruCache

class LazyBulk<T> private constructor(
    val totalRunSize: Int = 0,
    private val cachedSize: Int,
    private val onEvicted: ((index: Int, oldVal: T) -> Unit)? = null,
    private var onMiss: (Int) -> T = { throw IllegalStateException("onMiss not set") },
    private val lru: LRU<T>
) {
    constructor(
        totalRunSize: Int = 0,
        cachedSize: Int = 100,
        onEvicted: ((index: Int, oldVal: T) -> Unit)? = null,
        onMiss: (Int) -> T,
    ) : this(
        totalRunSize,
        cachedSize,
        onEvicted,
        onMiss,
        LRU(cachedSize, onEvicted)
    )

    private class LRU<T>(maxSize: Int, private val onEvicted: ((index: Int, oldVal: T) -> Unit)?) :
        LruCache<Int, T>(maxSize) {
        override fun entryRemoved(evicted: Boolean, key: Int, oldValue: T, newValue: T?) {
            if (onEvicted != null && evicted) onEvicted.invoke(key, oldValue)
        }
    }

    operator fun get(index: Int): T = lru.get(index) ?: onMiss.invoke(index)

    operator fun set(index: Int, value: T): T = lru.put(index, value)

    @Suppress("MemberVisibilityCanBePrivate")
    fun copy(
        totalRunSize: Int = this.totalRunSize,
        cachedSize: Int = this.cachedSize,
        onEvicted: ((index: Int, oldVal: T) -> Unit)? = this.onEvicted,
        onMiss: (Int) -> T = this.onMiss
    ) = LazyBulk(
        totalRunSize,
        cachedSize,
        onEvicted,
        onMiss,
        lru
    )

    fun copy(withAddedData: Collection<T>, totalRunSize: Int = this.totalRunSize) =
        copy(totalRunSize = totalRunSize).apply {
            putBulk(windowAround(totalRunSize, withAddedData.size), withAddedData)
        }

    private fun putBulk(index: IntRange, values: Collection<T>): Unit =
        values.forEachIndexed { i, value -> lru.put(index.first + i, value) }

    fun lookAhead(index: Int, peekSize: Int): Boolean {
        val peekRange = windowAround(index, peekSize)
        return lru.get(peekRange.first) != null && lru.get(peekRange.last) != null
    }

    private fun windowAround(index: Int, peekSize: Int) =
        ((index - peekSize)..(index + peekSize)).coerceInIndex()

    private fun IntRange.coerceInIndex() =
        first.coerceAtLeast(0)..last.coerceAtMost(totalRunSize - 1)
}