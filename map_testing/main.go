package main

import "fmt"
import "intset_testing"
import "iterator_testing"
import "os"
import "combined_testing"

func MillionOpsPerSecond(nGoroutines int, callback func(nGoroutines int) int64) float64 {
	nsOp := callback(nGoroutines)
	opS := float64(1000000000) / float64(nsOp)
	return opS / float64(1000000)
}

func main() {
	endEarly := false
	if endEarly {
		combinedSkimFile, err := os.Create("combinedSkim.csv")
		if err != nil {
			panic("Cannot create combinedSkim.csv")
		}

		// Header - CombinedSkim
		combinedSkimFile.WriteString(fmt.Sprintf("Map-CombinedSkim"))
		for i := 1; i <= 32; i = i << 1 {
			combinedSkimFile.WriteString(fmt.Sprintf(",%v", i))
		}
		combinedSkimFile.WriteString("\n")

		// // Concurrent Map
		// combinedSkimFile.WriteString(fmt.Sprintf("ConcurrentMap"))
		// for i := 1; i <= 32; i = i << 1 {
		// 	combinedSkimFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, combined_testing.ConcurrentCombinedSkim)))
		// }
		// combinedSkimFile.WriteString("\n")

		// Concurrent Map - Interlocked
		combinedSkimFile.WriteString(fmt.Sprintf("ConcurrentMap (Interlocked)"))
		for i := 1; i <= 32; i = i << 1 {
			combinedSkimFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, combined_testing.ConcurrentCombinedSkim_Interlocked)))
		}
		combinedSkimFile.WriteString("\n")

		// Synchronized Map
		combinedSkimFile.WriteString(fmt.Sprintf("SynchronizedMap (Mutex)"))
		for i := 1; i <= 32; i = i << 1 {
			combinedSkimFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, combined_testing.SynchronizedCombinedSkim)))
		}
		combinedSkimFile.WriteString("\n")

		// ReaderWriter Map
		combinedSkimFile.WriteString(fmt.Sprintf("ReaderWriterMap (RWMutex)"))
		for i := 1; i <= 32; i = i << 1 {
			combinedSkimFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, combined_testing.ReaderWriterCombinedSkim)))
		}
		combinedSkimFile.WriteString("\n")
		combinedSkimFile.Close()

		return
	}
	// Create files to dump information to
	var intsetFile, iteratorROFile, iteratorRWFile, combinedFile *os.File
	intsetFile, err := os.Create("intset.csv")
	if err != nil {
		panic("Cannot create intsetFile.csv")
	}
	iteratorROFile, err = os.Create("iteratorROFile.csv")
	if err != nil {
		panic("Cannot create iteratorROFile.csv")
	}
	iteratorRWFile, err = os.Create("iteratorRWFile.csv")
	if err != nil {
		panic("Cannot create iteratorRWFile.csv")
	}
	combinedFile, err = os.Create("combinedFile.csv")
	if err != nil {
		panic("Cannot create combinedFile.csv")
	}

	// Header - Intset
	intsetFile.WriteString(fmt.Sprintf("Map-Intset"))
	for i := 1; i <= 32; i = i << 1 {
		intsetFile.WriteString(fmt.Sprintf(",%v", i))
	}
	intsetFile.WriteString("\n")

	// Concurrent Map
	intsetFile.WriteString(fmt.Sprintf("ConcurrentMap"))
	for i := 1; i <= 32; i = i << 1 {
		intsetFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, intset_testing.ConcurrentIntset)))
	}
	intsetFile.WriteString("\n")

	// Synchronized Map
	intsetFile.WriteString(fmt.Sprintf("SynchronizedMap (Mutex)"))
	for i := 1; i <= 32; i = i << 1 {
		intsetFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, intset_testing.SynchronizedIntset)))
	}
	intsetFile.WriteString("\n")

	// ReaderWriterMap
	intsetFile.WriteString(fmt.Sprintf("ReaderWriterMap (RWMutex)"))
	for i := 1; i <= 32; i = i << 1 {
		intsetFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, intset_testing.ReaderWriterIntset)))
	}
	intsetFile.WriteString("\n")
	intsetFile.Close()

	// Header - IteratorRO
	iteratorROFile.WriteString(fmt.Sprintf("Map-IteratorRO"))
	for i := 1; i <= 32; i = i << 1 {
		iteratorROFile.WriteString(fmt.Sprintf(",%v", i))
	}
	iteratorROFile.WriteString("\n")

	// Concurrent Map
	// iteratorROFile.WriteString(fmt.Sprintf("ConcurrentMap"))
	// for i := 1; i <= 32; i = i << 1 {
	// 	iteratorROFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, iterator_testing.ConcurrentIterator_RO)))
	// }
	// iteratorROFile.WriteString("\n")

	// Concurrent Map - Interlocked
	iteratorROFile.WriteString(fmt.Sprintf("ConcurrentMap (Interlocked)"))
	for i := 1; i <= 32; i = i << 1 {
		iteratorROFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, iterator_testing.ConcurrentIterator_Interlocked_RO)))
	}
	iteratorROFile.WriteString("\n")

	// Default Map
	iteratorROFile.WriteString(fmt.Sprintf("DefaultMap (No Mutex)"))
	for i := 1; i <= 32; i = i << 1 {
		iteratorROFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, iterator_testing.DefaultIterator_RO)))
	}
	iteratorROFile.WriteString("\n")
	iteratorROFile.Close()

	// Header - IteratorRW
	iteratorRWFile.WriteString(fmt.Sprintf("Map-IteratorRW"))
	for i := 1; i <= 32; i = i << 1 {
		iteratorRWFile.WriteString(fmt.Sprintf(",%v", i))
	}
	iteratorRWFile.WriteString("\n")

	// Concurrent Map
	// iteratorRWFile.WriteString(fmt.Sprintf("ConcurrentMap"))
	// for i := 1; i <= 32; i = i << 1 {
	// 	iteratorRWFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, iterator_testing.ConcurrentIterator_RW)))
	// }
	// iteratorRWFile.WriteString("\n")

	// Concurrent Map - Interlocked
	iteratorRWFile.WriteString(fmt.Sprintf("ConcurrentMap (Interlocked)"))
	for i := 1; i <= 32; i = i << 1 {
		iteratorRWFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, iterator_testing.ConcurrentIterator_Interlocked_RW)))
	}
	iteratorRWFile.WriteString("\n")

	// Synchronized Map
	iteratorRWFile.WriteString(fmt.Sprintf("SynchronizedMap (Mutex)"))
	for i := 1; i <= 32; i = i << 1 {
		iteratorRWFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, iterator_testing.SynchronizedIterator_RW)))
	}
	iteratorRWFile.WriteString("\n")

	// ReaderWriter Map
	iteratorRWFile.WriteString(fmt.Sprintf("ReaderWriterMap (RWMutex)"))
	for i := 1; i <= 32; i = i << 1 {
		iteratorRWFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, iterator_testing.ReaderWriterIterator_RW)))
	}
	iteratorRWFile.WriteString("\n")
	iteratorRWFile.Close()

	// Header - Combined
	combinedFile.WriteString(fmt.Sprintf("Map-Combined"))
	for i := 1; i <= 32; i = i << 1 {
		combinedFile.WriteString(fmt.Sprintf(",%v", i))
	}
	combinedFile.WriteString("\n")

	// Concurrent Map
	// combinedFile.WriteString(fmt.Sprintf("ConcurrentMap"))
	// for i := 1; i <= 32; i = i << 1 {
	// 	combinedFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, combined_testing.ConcurrentCombined)))
	// }
	// combinedFile.WriteString("\n")

	// Concurrent Map - Interlocked
	combinedFile.WriteString(fmt.Sprintf("ConcurrentMap (Interlocked)"))
	for i := 1; i <= 32; i = i << 1 {
		combinedFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, combined_testing.ConcurrentCombined_Interlocked)))
	}
	combinedFile.WriteString("\n")

	// Synchronized Map
	combinedFile.WriteString(fmt.Sprintf("SynchronizedMap (Mutex)"))
	for i := 1; i <= 32; i = i << 1 {
		combinedFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, combined_testing.SynchronizedCombined)))
	}
	combinedFile.WriteString("\n")

	// ReaderWriter Map
	combinedFile.WriteString(fmt.Sprintf("ReaderWriterMap (RWMutex)"))
	for i := 1; i <= 32; i = i << 1 {
		combinedFile.WriteString(fmt.Sprintf(",%.2f", MillionOpsPerSecond(i, combined_testing.ReaderWriterCombined)))
	}
	combinedFile.WriteString("\n")
	combinedFile.Close()
}