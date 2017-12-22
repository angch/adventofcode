defmodule D04a do
  @moduledoc """
  Documentation for D04a.
  """

  def index do
    "input.txt"
    |> fileRead
    |> processFile
    |> main
  end

  @doc """
  Test cases

  ## Examples

      iex> D04a.processFile("aa bb cc dd ee") |> D04a.solve
      1
      iex> D04a.processFile("aa bb cc dd aa") |> D04a.solve
      0
      iex> D04a.processFile("aa bb cc dd aaa") |> D04a.solve
      1

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
  end

  def solve(input) do
    input
    |> Enum.filter(fn x ->
          Enum.uniq(x) == x
       end)
    |> length
  end

end

D04a.index
|> IO.inspect
