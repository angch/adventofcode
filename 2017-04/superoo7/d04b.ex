defmodule D04b do
  @moduledoc """
  Documentation for D04b.
  """
  def index do
    "input.txt"
    |> fileRead
    |> processFile
    |> main
  end

  @doc """
  Hello world.

  ## Examples

      iex> D04b.processFile("abcde fghij") |> D04b.solve
      1
      iex> D04b.processFile("abcde xyz ecdab") |> D04b.solve
      0
      iex> D04b.processFile("a ab abc abd abf abj") |> D04b.solve
      1
      iex> D04b.processFile("iiii oiii ooii oooi oooo") |> D04b.solve
      1
      iex> D04b.processFile("oiii ioii iioi iiio") |> D04b.solve
      0

  """
  def main(input) do
    input
    |> solve
  end

  def fileRead(filename) do
    {:ok, input} = File.read(filename)
    input
  end

  def processFile(input) do
    input
    |> String.split("\n")
    |> Enum.map(fn x -> String.split(x," ") end)
    |> Enum.map(fn x -> Enum.map(x, fn y -> String.split(y, "", trim: true) end) end)
    |> Enum.map(fn x -> Enum.map(x, fn y -> Enum.sort(y) end) end)
  end

  def solve(input) do
    input
    |>  Enum.filter(fn y ->
          Enum.uniq(y) == y
        end)
    |> length
  end
end

D04b.index
|> IO.inspect
